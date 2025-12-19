package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	todoController "github.com/jobpay/todo/internal/presentation/controller/todo"
)

func Setup(
	e *echo.Echo,
	showController *todoController.ShowController,
	listController *todoController.ListController,
	storeController *todoController.StoreController,
	updateController *todoController.UpdateController,
	deleteController *todoController.DeleteController,
) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// APIルート
	api := e.Group("/api")
	{
		todos := api.Group("/todos")
		{
			todos.GET("", listController.Handle)
			todos.GET("/:id", showController.Handle)
			todos.POST("", storeController.Handle)
			todos.PUT("/:id", updateController.Handle)
			todos.DELETE("/:id", deleteController.Handle)
		}
	}
}
