package example_usecase

import (
	"testing"

	"github.com/charmingruby/push/internal/domain/example/example_entity"
	"github.com/charmingruby/push/test/inmemory"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	exampleRepo    *inmemory.InMemoryExampleRepository
	exampleService *ExampleService
}

// initial setup
func (s *Suite) SetupSuite() {
	s.exampleRepo = inmemory.NewInMemoryExampleRepository()
	s.exampleService = NewExampleService(s.exampleRepo)
}

// executes before each test
func (s *Suite) SetupTest() {
	s.exampleRepo.Items = []example_entity.Example{}
}

// executes after each test
func (s *Suite) TearDownTest() {
	s.exampleRepo.Items = []example_entity.Example{}
}

// executes before each sub test
func (s *Suite) SetupSubTest() {
	s.exampleRepo.Items = []example_entity.Example{}
}

// executes after each sub test
func (s *Suite) TearDownSubTest() {
	s.exampleRepo.Items = []example_entity.Example{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
