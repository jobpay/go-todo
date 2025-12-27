package helper

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jobpay/todo/internal/infrastructure/di"
	"github.com/jobpay/todo/internal/presentation/controller"
	"github.com/jobpay/todo/internal/router"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	TestServerPort = ":18080"
	TestServerURL  = "http://localhost:18080"
)

func StartTestServer() error {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3307")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "todo_app_test")

	go func() {
		container := di.NewContainer()
		if err := container.Build(); err != nil {
			log.Fatalf("Failed to build DI container: %v", err)
		}

		err := container.Invoke(func(
			e *echo.Echo,
			controllers *controller.Controllers,
		) {
			router.Setup(e, controllers)

			e.HideBanner = true
			if err := e.Start(TestServerPort); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Failed to start test server: %v", err)
			}
		})

		if err != nil {
			log.Fatalf("Failed to invoke test server: %v", err)
		}
	}()

	time.Sleep(500 * time.Millisecond)
	return nil
}

func ConnectTestDB() (*gorm.DB, error) {
	dsn := "root:password@tcp(localhost:3307)/todo_app_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to test database: %w", err)
	}
	return db, nil
}

func CleanupTable(db *gorm.DB, tableName string) {
	if db != nil {
		db.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
	}
}

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}
