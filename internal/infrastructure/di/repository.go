package di

import (
	tagPersistence "github.com/jobpay/todo/internal/infrastructure/persistence/tag"
	todoPersistence "github.com/jobpay/todo/internal/infrastructure/persistence/todo"
)

func (c *Container) provideRepositories() error {
	repositories := []interface{}{
		todoPersistence.NewTodoRepository,
		tagPersistence.NewTagRepository,
	}

	for _, repo := range repositories {
		if err := c.container.Provide(repo); err != nil {
			return err
		}
	}

	return nil
}
