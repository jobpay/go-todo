package di

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/jobpay/todo/internal/domain/repository"
	"github.com/jobpay/todo/internal/infrastructure/database"
	todoPersistence "github.com/jobpay/todo/internal/infrastructure/persistence/todo"
	"go.uber.org/dig"
	"gorm.io/gorm"

	todoUseCase "github.com/jobpay/todo/internal/application/usecase/todo"
	todoController "github.com/jobpay/todo/internal/presentation/controller/todo"
)

type Container struct {
	container *dig.Container
}

func NewContainer() *Container {
	return &Container{
		container: dig.New(),
	}
}

func (c *Container) Build() error {
	if err := c.container.Provide(func() (*gorm.DB, error) {
		config := database.LoadConfigFromEnv()
		return database.NewMySQLDB(config)
	}); err != nil {
		return err
	}

	if err := c.container.Provide(func(db *gorm.DB) repository.TodoRepository {
		return todoPersistence.NewTodoRepository(db)
	}); err != nil {
		return err
	}

	if err := c.container.Provide(todoUseCase.NewShowUseCase); err != nil {
		return err
	}
	if err := c.container.Provide(todoUseCase.NewListUseCase); err != nil {
		return err
	}
	if err := c.container.Provide(todoUseCase.NewStoreUseCase); err != nil {
		return err
	}
	if err := c.container.Provide(todoUseCase.NewUpdateUseCase); err != nil {
		return err
	}
	if err := c.container.Provide(todoUseCase.NewDeleteUseCase); err != nil {
		return err
	}

	if err := c.container.Provide(todoController.NewShowController); err != nil {
		return err
	}
	if err := c.container.Provide(todoController.NewListController); err != nil {
		return err
	}
	if err := c.container.Provide(todoController.NewStoreController); err != nil {
		return err
	}
	if err := c.container.Provide(todoController.NewUpdateController); err != nil {
		return err
	}
	if err := c.container.Provide(todoController.NewDeleteController); err != nil {
		return err
	}

	if err := c.container.Provide(func() *echo.Echo {
		return echo.New()
	}); err != nil {
		return err
	}

	log.Println("DI container built successfully")
	return nil
}

func (c *Container) Invoke(function interface{}) error {
	return c.container.Invoke(function)
}
