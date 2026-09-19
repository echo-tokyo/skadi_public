package repository

import (
	"log"
	"os"
	"testing"

	"skadi/backend/internal/pkg/db"
)

var _testRepoDB *RepoDB // db repo instance

// Use env-variable TEST_DSN to set data source for db connection.
func TestMain(m *testing.M) {
	dbStorage, err := db.New(os.Getenv("TEST_DSN"),
		db.WithDebugLogLevel(),
		db.WithDisableColorful(),
		db.WithTranslateError(),
		db.WithIgnoreNotFound())
	if err != nil {
		log.Fatalf("connect to db: %v", err)
	}

	_testRepoDB = NewRepoDB(dbStorage)
	os.Exit(m.Run())
}

func TestGetOneFull(t *testing.T) {
	userID := 1

	userObj, err := _testRepoDB.GetOneFull("id", userID)
	if err != nil {
		log.Fatalf("get by id: %v", err)
	}
	t.Logf("Gotten user with profile: %+v", userObj)
}
