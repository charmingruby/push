package container

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	testDB := NewMongoTestDatabase()
	defer testDB.Teardown()
	os.Exit(m.Run())
}
