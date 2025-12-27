package main

import (
	"log"

	"github.com/jobpay/todo/internal/infrastructure/di"
	"github.com/jobpay/todo/internal/presentation/controller"
	"github.com/jobpay/todo/internal/router"
	"github.com/labstack/echo/v4"
)

func main() {
	// DIコンテナを構築
	container := di.NewContainer()
	if err := container.Build(); err != nil {
		log.Fatalf("Failed to build DI container: %v", err)
	}

	err := container.Invoke(func(
		e *echo.Echo,
		controllers *controller.Controllers,
	) {
		router.Setup(e, controllers)

		log.Println("Server starting on :8080")
		if err := e.Start(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Failed to invoke router setup: %v", err)
	}
}
