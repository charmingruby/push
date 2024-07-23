package integration

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/infra/transport/rest"
	v1 "github.com/charmingruby/push/internal/infra/transport/rest/endpoint/v1"
	"github.com/charmingruby/push/test/factory"
	"github.com/charmingruby/push/test/integration/helper"
	"github.com/oklog/ulid/v2"
)

func (s *Suite) Test_V1ScheduleNotificationEndpoint() {
	s.Run("it should be able to schedule a notification", func() {
		communicationChannel, err := factory.MakeCommunicationChannel(
			s.communicationChannelRepo,
			"Email",
			"Description...",
		)
		s.NoError(err)

		payload := v1.ScheduleNotificationRequest{
			Destination:            "dummy@email.com",
			RawDate:                "2024-07-10 15:00",
			CommunicationChannelID: communicationChannel.ID,
		}
		body, err := json.Marshal(payload)
		s.NoError(err)

		route := s.V1Route("/notifications")

		res, err := http.Post(route, contentType, helper.WriteBody(body))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusCreated, res.StatusCode)

		resultantData := rest.Response{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal("notification created successfully", resultantData.Message)
	})

	s.Run("it should be not able to schedule a notification with an invalid payload", func() {
		communicationChannel, err := factory.MakeCommunicationChannel(
			s.communicationChannelRepo,
			"Email",
			"Description...",
		)
		s.NoError(err)

		payload := v1.ScheduleNotificationRequest{
			Destination:            "",
			RawDate:                "2024-07-10 15:00",
			CommunicationChannelID: communicationChannel.ID,
		}
		body, err := json.Marshal(payload)
		s.NoError(err)

		route := s.V1Route("/notifications")

		res, err := http.Post(route, contentType, helper.WriteBody(body))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusBadRequest, res.StatusCode)
	})

	s.Run("it should be not able to schedule a notification with invalid params", func() {
		communicationChannel, err := factory.MakeCommunicationChannel(
			s.communicationChannelRepo,
			"Email",
			"Description...",
		)
		s.NoError(err)

		invalidRawDate := "2024-07-10 15"
		payload := v1.ScheduleNotificationRequest{
			Destination:            "dummy@email.com",
			RawDate:                invalidRawDate,
			CommunicationChannelID: communicationChannel.ID,
		}
		body, err := json.Marshal(payload)
		s.NoError(err)

		route := s.V1Route("/notifications")

		res, err := http.Post(route, contentType, helper.WriteBody(body))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusUnprocessableEntity, res.StatusCode)

		resultantData := rest.Response{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal(
			fmt.Sprintf("unable to parse `%s` into date format: `2006-01-02 15:00`", invalidRawDate),
			resultantData.Message,
		)
	})

	s.Run("it should be not able to schedule a notification if communication channel dont exists", func() {
		payload := v1.ScheduleNotificationRequest{
			Destination:            "dummy@email.com",
			RawDate:                "2024-07-10 15:00",
			CommunicationChannelID: ulid.Make().String(),
		}
		body, err := json.Marshal(payload)
		s.NoError(err)

		route := s.V1Route("/notifications")

		res, err := http.Post(route, contentType, helper.WriteBody(body))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusNotFound, res.StatusCode)

		resultantData := rest.Response{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal(core.NewNotFoundErr("communication channel").Error(), resultantData.Message)
	})
}
