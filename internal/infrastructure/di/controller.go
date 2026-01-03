package di

import (
	"github.com/jobpay/todo/internal/presentation/controller"
	tagController "github.com/jobpay/todo/internal/presentation/controller/tag"
	todoController "github.com/jobpay/todo/internal/presentation/controller/todo"
)

func (c *Container) provideControllers() error {
	controllers := []interface{}{
		todoController.NewShowController,
		todoController.NewListController,
		todoController.NewStoreController,
		todoController.NewUpdateController,
		todoController.NewDeleteController,
		todoController.NewControllers,
		tagController.NewShowController,
		tagController.NewListController,
		tagController.NewStoreController,
		tagController.NewUpdateController,
		tagController.NewDeleteController,
		tagController.NewControllers,
		controller.NewControllers,
	}

	for _, ctrl := range controllers {
		if err := c.container.Provide(ctrl); err != nil {
			return err
		}
	}

	return nil
}
