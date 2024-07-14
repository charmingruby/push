package fake

import (
	"fmt"

	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func NewFakeDispatcher() *FakeDispatcher {
	return &FakeDispatcher{}
}

func (d *FakeDispatcher) Notify(n *notification_entity.Notification) error {
	if n == nil {
		return fmt.Errorf("notification dispatch error")
	} // simulates an error

	return nil
}

type FakeDispatcher struct{}
