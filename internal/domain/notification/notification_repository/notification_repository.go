package notification_repository

import (
	"time"

	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

type NotificationRepository interface {
	Store(n *notification_entity.Notification) error
	FindByID(id string) (*notification_entity.Notification, error)
	ListAvailableNotificationsBeforeADate(date time.Time) ([]notification_entity.Notification, error)
	SaveStatus(n *notification_entity.Notification) error
}
