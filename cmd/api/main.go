package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/jobpay/todo/internal/infrastructure/di"
	"github.com/jobpay/todo/internal/presentation/controller/todo"
	"github.com/jobpay/todo/internal/router"
)

func main() {
	// DIコンテナを構築
	container := di.NewContainer()
	if err := container.Build(); err != nil {
		log.Fatalf("Failed to build DI container: %v", err)
	}

	// ルーターをセットアップ
	err := container.Invoke(func(
		e *echo.Echo,
		showController *todo.ShowController,
		listController *todo.ListController,
		storeController *todo.StoreController,
		updateController *todo.UpdateController,
		deleteController *todo.DeleteController,
	) {
		router.Setup(e, showController, listController, storeController, updateController, deleteController)

		// サーバー起動
		log.Println("Server starting on :8080")
		if err := e.Start(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Failed to invoke router setup: %v", err)
	}
}
