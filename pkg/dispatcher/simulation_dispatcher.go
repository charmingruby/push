package dispatcher

import (
	"fmt"
	"log/slog"
	"time"

	"math/rand"

	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
)

func NewSimulationDispatcher() *SimulationDispatcher {
	return &SimulationDispatcher{}
}

func (d *SimulationDispatcher) Notify(n *notification_entity.Notification) error {
	if err := d.simulateScenario(); err != nil {
		return err
	}

	slog.Info(
		fmt.Sprintf("[DISPATCHER] Notification#%s sent!", n.ID),
	) // simulation dispatch

	return nil
}

type SimulationDispatcher struct{}

func (d *SimulationDispatcher) simulateScenario() error {
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
