package container

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	testDB := NewMongoTestDatabase()
	defer testDB.TearDown()
	os.Exit(m.Run())
}
