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

func (r *InMemoryNotificationRepository) FindByID(id string) (*notification_entity.Notification, error) {
	for _, e := range r.Items {
		if e.ID == id {
			return &e, nil
		}
	}

	return nil, core.NewNotFoundErr("notification")
}

func (r *InMemoryNotificationRepository) ListAvailableNotificationsBeforeADate(date time.Time) ([]notification_entity.Notification, error) {
	filteredNotifications := []notification_entity.Notification{}

	for _, n := range r.Items {
		if n.Date.Before(date) && (n.Status == "PENDING" || n.Status == "RETRYING") {
			filteredNotifications = append(filteredNotifications, n)
		}
	}

	return filteredNotifications, nil
}

func (r *InMemoryNotificationRepository) SaveStatus(n *notification_entity.Notification) error {
	for idx, e := range r.Items {
		if e.ID == n.ID {
			if e.Retries != n.Retries {
				r.Items[idx].Retries = n.Retries
			}

			r.Items[idx].Status = n.Status

			return nil
		}
	}

	return core.NewNotFoundErr("notification")
}
