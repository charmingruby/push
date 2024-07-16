package notification_usecase

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func (s *Suite) Test_GetNotificationUseCase() {
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
	s.NoError(err)

	s.Run("it should be able to get a notification", func() {
		err := s.notificationRepo.Store(baseNotification)
		s.NoError(err)
		s.Equal(1, len(s.notificationRepo.Items))

		dto := notification_dto.GetNotificationDTO{
			NotificationID: baseNotification.ID,
		}

		notification, err := s.useCase.GetNotificationUseCase(dto)

		s.NoError(err)
		s.Equal(baseNotification.ID, notification.ID)
	})

	s.Run("it should be not able to get a notification with invalid id", func() {
		dto := notification_dto.GetNotificationDTO{
			NotificationID: "invalid id",
		}

		notification, err := s.useCase.GetNotificationUseCase(dto)

		s.Error(err)
		s.Nil(notification)
		s.Equal(core.NewNotFoundErr("notification").Error(), err.Error())
	})
}
