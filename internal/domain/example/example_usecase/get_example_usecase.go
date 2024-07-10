package example_usecase

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/example/example_entity"
)

func (s *ExampleService) GetExample(id string) (*example_entity.Example, error) {
	example, err := s.exampleRepo.FindByID(id)
	if err != nil {
		return nil, core.NewNotFoundErr("example")
	}

	return example, nil
}
