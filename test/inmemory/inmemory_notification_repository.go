package inmemory

import (
	"time"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func NewInMemoryNotificationRepository() *InMemoryNotificationRepository {
	return &InMemoryNotificationRepository{
		Items: []notification_entity.Notification{},
	}
}

type InMemoryNotificationRepository struct {
	Items []notification_entity.Notification
}

func (r *InMemoryNotificationRepository) Store(n *notification_entity.Notification) error {
	r.Items = append(r.Items, *n)
	return nil
}

func (r *InMemoryNotificationRepository) GetNotificationByID(id string) (*notification_entity.Notification, error) {
	for _, e := range r.Items {
		if e.ID == id {
			return &e, nil
		}
	}

	return nil, core.NewNotFoundErr("notification")
}

func (r *InMemoryNotificationRepository) ListAvailableNotificationsBeforeDate(date time.Time) ([]notification_entity.Notification, error) {
	filteredNotifications := []notification_entity.Notification{}

	for _, n := range r.Items {
		if n.Date.Before(date) && (n.Status == "PENDING" || n.Status == "RETRYING") {
			filteredNotifications = append(filteredNotifications, n)
		}
	}

	return filteredNotifications, nil
}

func (r *InMemoryNotificationRepository) SaveNotificationStatus(n *notification_entity.Notification) error {
	var idx int

	for cIdx, i := range r.Items {
		if i.ID == n.ID {
			idx = cIdx
		}
	}

	r.Items = append(r.Items[:idx], r.Items[idx+1])

	return nil
}
