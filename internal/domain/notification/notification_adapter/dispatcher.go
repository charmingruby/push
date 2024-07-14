package notification_adapter

import "github.com/charmingruby/push/internal/domain/notification/notification_entity"

type Dispatcher interface {
	Notify(n *notification_entity.Notification) error
}
