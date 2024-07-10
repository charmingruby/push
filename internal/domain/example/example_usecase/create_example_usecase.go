package example_usecase

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/example/example_dto"
	"github.com/charmingruby/push/internal/domain/example/example_entity"
)

func (s *ExampleService) CreateExample(dto example_dto.CreateExampleDTO) error {
	example, err := example_entity.NewExample(dto.Name)
	if err != nil {
		return err
	}

	if err := s.exampleRepo.Store(example); err != nil {
		return core.NewInternalErr("create example store")
	}

	return nil
}
