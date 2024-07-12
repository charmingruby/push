package notification_repository

import (
	"time"

	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

type NotificationRepository interface {
	Store(n *notification_entity.Notification) error
	GetNotificationByID(id string) (*notification_entity.Notification, error)
	ListAvailableNotificationsBeforeDate(date time.Time) ([]notification_entity.Notification, error)
	SaveNotificationStatus(n *notification_entity.Notification) error
}
