package integration

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/push/internal/core"
	v1 "github.com/charmingruby/push/internal/infra/transport/rest/endpoint/v1"
	"github.com/charmingruby/push/test/factory"
	"github.com/charmingruby/push/test/integration/helper"
	"github.com/oklog/ulid/v2"
)

func (s *Suite) Test_V1GetNotificationEndpoint() {

	s.Run("it should be able to get a notification", func() {
		communicationChannel, err := factory.MakeCommunicationChannel(
			s.communicationChannelRepo,
			"Email",
			"Email channel",
		)
		s.NoError(err)

		notification, err := factory.MakeNotification(
			s.notificationRepo,
			"dummy@email.com",
			"2024-07-10 15:00",
			communicationChannel.ID,
		)
		s.NoError(err)

		route := s.V1Route(
			fmt.Sprintf("/notifications/%s", notification.ID),
		)

		res, err := http.Get(route)
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusOK, res.StatusCode)

		resultantData := v1.GetNotificationResponse{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal("notification found", resultantData.Message)
		s.Equal(notification.ID, resultantData.Data.ID)
	})

	s.Run("it should be not able to get a nonexistent notification", func() {
		route := s.V1Route(
			fmt.Sprintf("/notifications/%s", ulid.Make().String()),
		)

		res, err := http.Get(route)
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusNotFound, res.StatusCode)

		resultantData := v1.GetNotificationResponse{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal(core.NewNotFoundErr("notification").Error(), resultantData.Message)
	})
}
