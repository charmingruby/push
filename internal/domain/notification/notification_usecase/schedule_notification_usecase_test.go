package notification_usecase

import (
	"fmt"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func (s *Suite) Test_ScheduleNotificationUseCase() {
	destination := "dummy@email.com"
	rawDate := "2024-07-10 15:00"
	communicationChannel, err := notification_entity.NewCommunicationChannel(
		"dummy communication",
		"dummy description",
	)
	s.NoError(err)

	s.Run("it should be able to schedule a notification successfully", func() {
		err := s.communicationChannelRepo.Store(communicationChannel)
		s.NoError(err)
		s.Equal(len(s.communicationChannelRepo.Items), 1)

		dto := notification_dto.ScheduleNotificationDTO{
			Destination:            destination,
			RawDate:                rawDate,
			CommunicationChannelID: communicationChannel.ID,
		}

		err = s.useCase.ScheduleNotificationUseCase(dto)

		s.NoError(err)
		s.Equal(len(s.notificationRepo.Items), 1)
	})

	s.Run("it should be not able to schedule a notification with a invalid communication channel id ", func() {
		dto := notification_dto.ScheduleNotificationDTO{
			Destination:            destination,
			RawDate:                rawDate,
			CommunicationChannelID: "invalid id",
		}

		err = s.useCase.ScheduleNotificationUseCase(dto)

		s.Error(err)
		s.Equal(core.NewNotFoundErr("communication channel").Error(), err.Error())
	})

	s.Run("it should be not able to schedule a notification with entity error", func() {
		refDate := "2006-01-02 15:00"
		invalidRawDate := "2024-07-10 15"

		err := s.communicationChannelRepo.Store(communicationChannel)
		s.NoError(err)
		s.Equal(len(s.communicationChannelRepo.Items), 1)

		dto := notification_dto.ScheduleNotificationDTO{
			Destination:            destination,
			RawDate:                invalidRawDate,
			CommunicationChannelID: communicationChannel.ID,
		}

		err = s.useCase.ScheduleNotificationUseCase(dto)

		s.Error(err)
		s.Equal(core.NewValidationErr(
			fmt.Sprintf("unable to parse `%s` into date format: `%s`", invalidRawDate, refDate),
		).Error(), err.Error())
	})
}
