package integration

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/charmingruby/push/internal/domain/notification/notification_repository"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/charmingruby/push/internal/infra/database/mongo_repository"
	"github.com/charmingruby/push/internal/infra/observability/metric"
	"github.com/charmingruby/push/internal/infra/transport/rest"
	v1 "github.com/charmingruby/push/internal/infra/transport/rest/endpoint/v1"
	"github.com/charmingruby/push/pkg/dispatcher"
	"github.com/charmingruby/push/test/container"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

const (
	contentType = "application/json"
)

type Suite struct {
	suite.Suite
	container                *container.MongoTestDatabase
	server                   *httptest.Server
	handler                  *v1.HTTPHandler
	dispatcher               *dispatcher.SimulationDispatcher
	notificationRepo         notification_repository.NotificationRepository
	communicationChannelRepo notification_repository.CommunicationChannelRepository
}

func (s *Suite) SetupSuite() {
	tdb := container.NewMongoTestDatabase()
	s.container = tdb
}

func (s *Suite) TearDownSuite() {
	s.container.DB.Client().Disconnect(context.Background())
}

func (s *Suite) SetupSubTest() {
	err := s.container.CreateCollections()
	s.NoError(err)

	s.notificationRepo = mongo_repository.NewNotificationsMongoRepository(s.container.DB)
	s.communicationChannelRepo = mongo_repository.NewCommunicationChannelMongoRepository(s.container.DB)

	s.dispatcher = dispatcher.NewSimulationDispatcher()

	notificationSvc := notification_usecase.NewNotificationUseCaseRegistry(
		s.notificationRepo,
		s.communicationChannelRepo,
		s.dispatcher,
	)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	metrics := metric.NewMetrics()

	s.handler = v1.NewHTTPHandler(router, notificationSvc, metrics)
	s.handler.Register()
	server := rest.NewServer(router, "3000")
	s.server = httptest.NewServer(server.Router)
}

func (s *Suite) TearDownSubTest() {
	err := s.container.DropCollections()
	s.NoError(err)
	s.server.Close()
}

func (s *Suite) V1Route(path string) string {
	return fmt.Sprintf("%s/api/v1%s", s.server.URL, path)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
