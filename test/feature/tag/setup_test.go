package tag_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/jobpay/todo/test/feature/helper"
	"gorm.io/gorm"
)

var (
	testDB        *gorm.DB
	httpClient    *http.Client
	testServerURL = helper.TestServerURL
)

func TestMain(m *testing.M) {
	if err := setupTestEnvironment(); err != nil {
		log.Fatalf("Failed to setup test environment: %v", err)
	}

	m.Run()
}

func setupTestEnvironment() error {
	if err := helper.StartTestServer(); err != nil {
		return err
	}

	db, err := helper.ConnectTestDB()
	if err != nil {
		return err
	}
	testDB = db

	httpClient = helper.NewHTTPClient()

	return nil
}

func cleanupDB() {
	helper.CleanupTable(testDB, "tags")
}

