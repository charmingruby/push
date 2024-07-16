package notification_usecase

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func (s *Suite) Test_CreateCommunicationChannelUseCase() {
	name := "dummy name"
	description := "dummy description"

	s.Run("it should be able to create communication channel successfully", func() {
		dto := notification_dto.CreateCommunicationChannelDTO{
			Name:        name,
			Description: description,
		}

		err := s.useCase.CreateCommunicationChannelUseCase(dto)

		s.NoError(err)
		s.Equal(len(s.communicationChannelRepo.Items), 1)
		s.Equal(s.communicationChannelRepo.Items[0].Name, name)
	})

	s.Run("it should be not able to create communication channel with conflicting name", func() {
		communicationChannel, err := notification_entity.NewCommunicationChannel(
			name,
			description,
		)
		s.NoError(err)
		s.communicationChannelRepo.Items = append(s.communicationChannelRepo.Items, *communicationChannel)
		s.Equal(s.communicationChannelRepo.Items[0].Name, name)
		s.Equal(len(s.communicationChannelRepo.Items), 1)

		dto := notification_dto.CreateCommunicationChannelDTO{
			Name:        name,
			Description: description,
		}

		err = s.useCase.CreateCommunicationChannelUseCase(dto)

		s.Error(err)
		s.Equal(core.NewConflictErr("communication channel", "name").Error(), err.Error())
	})

	s.Run("it should be not able to create communication channel with entity error", func() {
		dto := notification_dto.CreateCommunicationChannelDTO{
			Name:        "",
			Description: description,
		}

		err := s.useCase.CreateCommunicationChannelUseCase(dto)

		s.Error(err)
		s.Equal(core.NewValidationErr(core.ErrRequired("name")).Error(), err.Error())
		s.Equal(len(s.communicationChannelRepo.Items), 0)
	})
}
