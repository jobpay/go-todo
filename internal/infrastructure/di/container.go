package di

import (
	"log"

	"go.uber.org/dig"
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
	if err := c.provideInfrastructure(); err != nil {
		return err
	}

	if err := c.provideRepositories(); err != nil {
		return err
	}

	if err := c.provideUseCases(); err != nil {
		return err
	}

	if err := c.provideControllers(); err != nil {
		return err
	}

	log.Println("DI container built successfully")
	return nil
}

func (c *Container) Invoke(function interface{}) error {
	return c.container.Invoke(function)
}
