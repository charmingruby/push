package factory

import (
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	"github.com/charmingruby/push/internal/domain/notification/notification_repository"
)

func MakeNotification(
	repo notification_repository.NotificationRepository,
	destination,
	rawDate,
	communicationChannelID string,
) (*notification_entity.Notification, error) {
	n, err := notification_entity.NewNotification(
		destination,
		rawDate,
		communicationChannelID,
	)
	if err != nil {
		return nil, err
	}

	if err := repo.Store(n); err != nil {
		return nil, err
	}

	return n, nil
}
