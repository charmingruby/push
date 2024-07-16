package notification_usecase

import (
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func (s *Suite) Test_CheckAndSendNotificationUseCase() {
	communicationChannel, err := notification_entity.NewCommunicationChannel(
		"dummy communication",
		"dummy description",
	)
	s.NoError(err)
	err = s.communicationChannelRepo.Store(communicationChannel)
	s.NoError(err)

	destination := "dummy destination"
	rawDate := "2006-01-02 15:00"

	s.Run("it should be able to check and send notifications", func() {
		notification, err := notification_entity.NewNotification(
			destination, rawDate, communicationChannel.ID,
		)
		s.NoError(err)
		s.Equal("PENDING", notification.Status)

		err = s.notificationRepo.Store(notification)
		s.NoError(err)
		s.Equal(1, len(s.notificationRepo.Items))

		notificationsWithFailures, err := s.useCase.CheckAndSendNotificationUseCase()

		s.NoError(err)
		s.Equal(0, len(notificationsWithFailures))
		s.Equal(1, len(s.notificationRepo.Items))
		s.Equal("SENT", s.notificationRepo.Items[0].Status)
		s.Equal(notification.ID, s.notificationRepo.Items[0].ID)
	})

	s.Run("it should be able to trigger a notification retry", func() {
		notificationToRetry, err := notification_entity.NewNotification(
			"trigger retry", rawDate, communicationChannel.ID,
		)
		s.NoError(err)
		s.Equal("PENDING", notificationToRetry.Status)
		s.Equal(0, notificationToRetry.Retries)

		err = s.notificationRepo.Store(notificationToRetry)
		s.NoError(err)
		s.Equal(1, len(s.notificationRepo.Items))

		notificationsWithFailures, err := s.useCase.CheckAndSendNotificationUseCase()

		s.NoError(err)
		s.Equal(0, len(notificationsWithFailures))
		s.Equal(1, len(s.notificationRepo.Items))
		s.Equal(notificationToRetry.ID, s.notificationRepo.Items[0].ID)
		s.Equal("RETRYING", s.notificationRepo.Items[0].Status)
		s.Equal(1, s.notificationRepo.Items[0].Retries)
	})

	s.Run("it should be not able to retry a notification with max retries quantity", func() {
		notificationToRetry, err := notification_entity.NewNotification(
			"trigger retry", rawDate, communicationChannel.ID,
		)
		s.NoError(err)
		notificationToRetry.Status = "RETRYING"
		notificationToRetry.Retries = notification_entity.MAX_RETRIES

		err = s.notificationRepo.Store(notificationToRetry)
		s.NoError(err)
		s.Equal(1, len(s.notificationRepo.Items))

		notificationsWithFailures, err := s.useCase.CheckAndSendNotificationUseCase()

		s.NoError(err)
		s.Equal(0, len(notificationsWithFailures))
		s.Equal(1, len(s.notificationRepo.Items))
		s.Equal(notificationToRetry.ID, s.notificationRepo.Items[0].ID)
		s.Equal("FAILURE", s.notificationRepo.Items[0].Status)
		s.Equal(notification_entity.MAX_RETRIES, s.notificationRepo.Items[0].Retries)
	})
}
