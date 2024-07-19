package integration

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	v1 "github.com/charmingruby/push/internal/infra/transport/rest/endpoint/v1"
	"github.com/charmingruby/push/test/factory"
	"github.com/charmingruby/push/test/integration/helper"
	"github.com/oklog/ulid/v2"
)

func (s *Suite) Test_V1CancelNotificationEndpoint() {
	s.Run("it should be able to cancel a notification", func() {
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
			fmt.Sprintf("/notifications/%s/cancel", notification.ID),
		)

		req, err := http.NewRequest(http.MethodPatch, route, nil)
		s.NoError(err)

		client := &http.Client{}
		res, err := client.Do(req)
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusOK, res.StatusCode)

		resultantData := v1.Response{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal("notification canceled successfully", resultantData.Message)

		notificationFound, err := s.notificationRepo.FindByID(notification.ID)
		s.NoError(err)

		s.Equal(notification.ID, notificationFound.ID)
		s.Equal("CANCELED", notificationFound.Status)
	})

	s.Run("it should be not able to cancel a notification already canceled", func() {
		communicationChannel, err := factory.MakeCommunicationChannel(
			s.communicationChannelRepo,
			"Email",
			"Email channel",
		)
		s.NoError(err)

		notification, err := notification_entity.NewNotification(
			"dummy@email.com",
			"2024-07-10 15:00",
			communicationChannel.ID,
		)
		s.NoError(err)

		notification.StatusCanceled()

		err = s.notificationRepo.Store(notification)
		s.NoError(err)

		route := s.V1Route(
			fmt.Sprintf("/notifications/%s/cancel", notification.ID),
		)

		req, err := http.NewRequest(http.MethodPatch, route, nil)
		s.NoError(err)

		client := &http.Client{}
		res, err := client.Do(req)
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusUnprocessableEntity, res.StatusCode)

		resultantData := v1.Response{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal("notification is already canceled", resultantData.Message)
	})

	s.Run("it should be not able to cancel a notification already sent", func() {
		communicationChannel, err := factory.MakeCommunicationChannel(
			s.communicationChannelRepo,
			"Email",
			"Email channel",
		)
		s.NoError(err)

		notification, err := notification_entity.NewNotification(
			"dummy@email.com",
			"2024-07-10 15:00",
			communicationChannel.ID,
		)
		s.NoError(err)

		notification.StatusSent()

		err = s.notificationRepo.Store(notification)
		s.NoError(err)

		route := s.V1Route(
			fmt.Sprintf("/notifications/%s/cancel", notification.ID),
		)

		req, err := http.NewRequest(http.MethodPatch, route, nil)
		s.NoError(err)

		client := &http.Client{}
		res, err := client.Do(req)
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusUnprocessableEntity, res.StatusCode)

		resultantData := v1.Response{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal("notification is already sent", resultantData.Message)
	})

	s.Run("it should be not able to cancel a nonexistent notification", func() {
		route := s.V1Route(
			fmt.Sprintf("/notifications/%s/cancel", ulid.Make().String()),
		)

		req, err := http.NewRequest(http.MethodPatch, route, nil)
		s.NoError(err)

		client := &http.Client{}
		res, err := client.Do(req)
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusNotFound, res.StatusCode)

		resultantData := v1.Response{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal(core.NewNotFoundErr("notification").Error(), resultantData.Message)
	})
}
