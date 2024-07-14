package fake

import (
	"fmt"
	"log/slog"
	"time"

	"math/rand"

	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func NewFakeDispatcher() *FakeDispatcher {
	return &FakeDispatcher{}
}

func (d *FakeDispatcher) Notify(n *notification_entity.Notification) error {
	if err := d.simulateScenario(); err != nil {
		return err
	}

	slog.Info(
		fmt.Sprintf("[DISPATCHER] Notification#%s sent!", n.ID),
	) // fake dispatch

	return nil
}

type FakeDispatcher struct{}

func (d *FakeDispatcher) simulateScenario() error {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	min := 1
	max := 3
	randomInt := rng.Intn(max-min+1) + min

	if randomInt%2 == 0 {
		return fmt.Errorf("notification dispatch error")
	}

	return nil
}
