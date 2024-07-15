package notification_usecase

import (
	"testing"

	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	"github.com/charmingruby/push/test/fake"
	"github.com/charmingruby/push/test/inmemory"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	communicationChannelRepo *inmemory.InMemoryCommunicationChannelRepository
	notificationRepo         *inmemory.InMemoryNotificationRepository
	notificationUc           *NotificationUseCaseRegistry
}

func (s *Suite) SetupSuite() {
	s.communicationChannelRepo = inmemory.NewInMemoryCommunicationChannelRepository()
	s.notificationRepo = inmemory.NewInMemoryNotificationRepository()

	fakeDispatcher := fake.NewFakeDispatcher()

	s.notificationUc = NewNotificationUseCaseRegistry(s.notificationRepo, s.communicationChannelRepo, fakeDispatcher)
}

func (s *Suite) SetupTest() {
	s.communicationChannelRepo.Items = []notification_entity.CommunicationChannel{}
	s.notificationRepo.Items = []notification_entity.Notification{}
}

func (s *Suite) TearDownTest() {
	s.communicationChannelRepo.Items = []notification_entity.CommunicationChannel{}
	s.notificationRepo.Items = []notification_entity.Notification{}
}

func (s *Suite) SetupSubTest() {
	s.communicationChannelRepo.Items = []notification_entity.CommunicationChannel{}
	s.notificationRepo.Items = []notification_entity.Notification{}
}

func (s *Suite) TearDownSubTest() {
	s.communicationChannelRepo.Items = []notification_entity.CommunicationChannel{}
	s.notificationRepo.Items = []notification_entity.Notification{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
