package example_entity

import (
	"time"

	"github.com/charmingruby/push/internal/core"
	"github.com/google/uuid"
)

func NewExample(name string) (*Example, error) {
	e := Example{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	if err := core.ValidateStruct(e); err != nil {
		return nil, err
	}

	return &e, nil
}

type Example struct {
	ID        string    `json:"id" validate:"required" db:"id"`
	Name      string    `json:"name" validate:"min=3,max=16" db:"name"`
	CreatedAt time.Time `json:"created_at" validate:"required" db:"created_at"`
}
