package di

import (
	todoPersistence "github.com/jobpay/todo/internal/infrastructure/persistence/todo"
)

func (c *Container) provideRepositories() error {
	return c.container.Provide(todoPersistence.NewTodoRepository)
}
