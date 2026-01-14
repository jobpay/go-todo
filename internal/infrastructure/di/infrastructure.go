package di

import (
	"github.com/jobpay/todo/internal/infrastructure/database"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (c *Container) provideInfrastructure() error {
	// Database
	if err := c.container.Provide(func() (*gorm.DB, error) {
		config := database.LoadConfigFromEnv()
		return database.NewMySQLDB(config)
	}); err != nil {
		return err
	}

	// Redis
	if err := c.container.Provide(func() (*redis.Client, error) {
		config := database.LoadRedisConfigFromEnv()
		return database.NewRedisClient(config)
	}); err != nil {
		return err
	}

	// Web Framework
	if err := c.container.Provide(func() *echo.Echo {
		return echo.New()
	}); err != nil {
		return err
	}

	return nil
}
