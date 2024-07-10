package inmemory

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/example/example_entity"
)

func NewInMemoryExampleRepository() *InMemoryExampleRepository {
	return &InMemoryExampleRepository{
		Items: []example_entity.Example{},
	}
}

type InMemoryExampleRepository struct {
	Items []example_entity.Example
}

func (r *InMemoryExampleRepository) Store(e *example_entity.Example) error {
	r.Items = append(r.Items, *e)
	return nil
}

func (r *InMemoryExampleRepository) FindByID(id string) (*example_entity.Example, error) {
	for _, e := range r.Items {
		if e.ID == id {
			return &e, nil
		}
	}

	return nil, core.NewNotFoundErr("example")
}
