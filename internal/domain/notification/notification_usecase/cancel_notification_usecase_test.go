package notification_usecase

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func (s *Suite) Test_CancelNotificationUseCase() {
	communicationChannel, err := notification_entity.NewCommunicationChannel(
		"dummy communication",
		"dummy description",
	)
	s.NoError(err)
	err = s.communicationChannelRepo.Store(communicationChannel)
	s.NoError(err)

	baseNotification, err := notification_entity.NewNotification(
		"dummy@email.com",
		"2024-07-10 15:00",
		communicationChannel.ID,
	)

	s.Run("it should be able to cancel a notification", func() {
		err = s.notificationRepo.Store(baseNotification)
		s.NoError(err)
		s.Equal("PENDING", s.notificationRepo.Items[0].Status)

		dto := notification_dto.CancelNotificationDTO{
			NotificationID: baseNotification.ID,
		}

		err = s.useCase.CancelNotiticationUseCase(dto)

		s.NoError(err)
		s.Equal("CANCELED", s.notificationRepo.Items[0].Status)
	})

	s.Run("it should be not able to cancel a notification if is not found", func() {
		dto := notification_dto.CancelNotificationDTO{
			NotificationID: baseNotification.ID,
		}

		err = s.useCase.CancelNotiticationUseCase(dto)

		s.Error(err)
		s.Equal(core.NewNotFoundErr("notification").Error(), err.Error())
		s.Equal(0, len(s.notificationRepo.Items))
	})

	s.Run("it should be able to cancel a notification if is already sent", func() {
		sentNotification := *baseNotification
		sentNotification.StatusSent()

		err = s.notificationRepo.Store(&sentNotification)
		s.NoError(err)
		s.Equal("SENT", s.notificationRepo.Items[0].Status)

		dto := notification_dto.CancelNotificationDTO{
			NotificationID: sentNotification.ID,
		}

		err = s.useCase.CancelNotiticationUseCase(dto)

		s.Error(err)
		s.Equal(core.NewValidationErr("notification is already sent").Error(), err.Error())
	})
}
