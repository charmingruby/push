package fake

import (
	"fmt"

	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func NewFakeDispatcher() *FakeDispatcher {
	return &FakeDispatcher{}
}

func (d *FakeDispatcher) Notify(n *notification_entity.Notification) error {
	if n.Destination == "trigger retry" {
		return fmt.Errorf("notification dispatch error")
	} // simulates an error

	d.NotificationsSent = append(d.NotificationsSent, *n)

	return nil
}

type FakeDispatcher struct {
	NotificationsSent []notification_entity.Notification
}
