package notification_entity

import (
	"time"

	"github.com/charmingruby/push/internal/core"
	"github.com/oklog/ulid/v2"
)

func NewCommunicationChannel(name, description string) (*CommunicationChannel, error) {
	e := CommunicationChannel{
		ID:          ulid.Make().String(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}

	if err := core.ValidateStruct(e); err != nil {
		return nil, err
	}

	return &e, nil
}

type CommunicationChannel struct {
	ID          string    `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at" validate:"required"`
}
