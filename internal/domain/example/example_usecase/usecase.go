package example_usecase

import (
	"github.com/charmingruby/push/internal/domain/example/example_dto"
	"github.com/charmingruby/push/internal/domain/example/example_entity"
	"github.com/charmingruby/push/internal/domain/example/example_repository"
)

type ExampleServiceContract interface {
	CreateExample(dto example_dto.CreateExampleDTO) error
	GetExample(id string) (*example_entity.Example, error)
}

func NewExampleService(exampleRepo example_repository.ExampleRepository) *ExampleService {
	return &ExampleService{
		exampleRepo: exampleRepo,
	}
}

type ExampleService struct {
	exampleRepo example_repository.ExampleRepository
}
