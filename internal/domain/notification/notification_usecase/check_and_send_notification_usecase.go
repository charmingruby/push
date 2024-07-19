package notification_usecase

import (
	"time"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func (r *NotificationUseCaseRegistry) CheckAndSendNotificationUseCase() ([]notification_entity.Notification, error) {
	now := time.Now()

	notificationsToBeSend, err := r.notificationRepo.ListAvailableNotificationsBeforeADate(now) // STATUS: pending or retrying
	if err != nil {
		return nil, core.NewInternalErr(
			"check and send notifications use case: list available notifications before date",
		)
	}

	var notificationsWithInternalErr []notification_entity.Notification

	if len(notificationsToBeSend) != 0 {
		for _, n := range notificationsToBeSend {
			if n.Status == "RETRYING" {
				err := n.Retry()
				if err != nil {
					n.StatusFailure()
					if err := r.notificationRepo.SaveStatus(&n); err != nil {
						notificationsWithInternalErr = append(notificationsWithInternalErr, n)
					}

					continue
				}
			}

			if err := r.dispatcher.Notify(&n); err != nil {
				if n.Status != "RETRYING" {
					n.StatusRetrying()
					n.Retry()
				}

				if err := r.notificationRepo.SaveStatus(&n); err != nil {
					notificationsWithInternalErr = append(notificationsWithInternalErr, n)
				}

				continue
			}

			n.StatusSent()
			if err := r.notificationRepo.SaveStatus(&n); err != nil {
				notificationsWithInternalErr = append(notificationsWithInternalErr, n)
			}
		}
	}

	return notificationsWithInternalErr, nil
}
