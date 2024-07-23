package integration

import (
	"encoding/json"
	"net/http"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/infra/transport/rest"
	v1 "github.com/charmingruby/push/internal/infra/transport/rest/endpoint/v1"
	"github.com/charmingruby/push/test/factory"
	"github.com/charmingruby/push/test/integration/helper"
)

func (s *Suite) Test_V1CreateCommunicationChannelEndpoint() {
	s.Run("it should be able to create a communication channel successfully", func() {
		payload := v1.CreateCommunicationChannelRequest{
			Name:        "Email",
			Description: "Email channel",
		}
		body, err := json.Marshal(payload)
		s.NoError(err)

		route := s.V1Route("/communication-channels")

		res, err := http.Post(route, contentType, helper.WriteBody(body))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusCreated, res.StatusCode)

		resultantData := rest.Response{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal("communication channel created successfully", resultantData.Message)
	})

	s.Run("it should be not able to create communication channel with an invalid payload", func() {
		payload := v1.CreateCommunicationChannelRequest{
			Name: "Email",
		}
		body, err := json.Marshal(payload)
		s.NoError(err)

		route := s.V1Route("/communication-channels")

		res, err := http.Post(route, contentType, helper.WriteBody(body))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusBadRequest, res.StatusCode)
	})

	s.Run("it should be not able to create communication channel with invalid params", func() {
		payload := v1.CreateCommunicationChannelRequest{
			Name:        "Email",
			Description: "Ema",
		}
		body, err := json.Marshal(payload)
		s.NoError(err)

		route := s.V1Route("/communication-channels")

		res, err := http.Post(route, contentType, helper.WriteBody(body))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusUnprocessableEntity, res.StatusCode)

		resultantData := rest.Response{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal(core.ErrMinLength("description", "4"), resultantData.Message)
	})

	s.Run("it should be not able to create communication channel with a conflicting name", func() {
		conflictingName := "Email"
		description := "dummy description"

		conflictingCommunicationChannel, err := factory.MakeCommunicationChannel(
			s.communicationChannelRepo,
			conflictingName,
			description,
		)
		s.NoError(err)

		conflictingCommunicationChannelFound, err := s.communicationChannelRepo.FindByName(conflictingName)
		s.NoError(err)
		s.Equal(conflictingCommunicationChannel.ID, conflictingCommunicationChannelFound.ID)

		payload := v1.CreateCommunicationChannelRequest{
			Name:        conflictingName,
			Description: description,
		}
		body, err := json.Marshal(payload)
		s.NoError(err)

		route := s.V1Route("/communication-channels")

		res, err := http.Post(route, contentType, helper.WriteBody(body))
		s.NoError(err)
		defer res.Body.Close()

		s.Equal(http.StatusConflict, res.StatusCode)

		resultantData := rest.Response{}
		err = helper.ParseRequest(&resultantData, res.Body)
		s.NoError(err)

		s.Equal(core.NewConflictErr("communication channel", "name").Error(), resultantData.Message)
	})
}
