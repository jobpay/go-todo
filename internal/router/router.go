package router

import (
	"github.com/jobpay/todo/internal/presentation/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo, controllers *controller.Controllers) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	api := e.Group("/api")
	{
		todos := api.Group("/todos")
		{
			todos.GET("", controllers.Todo.List.Handle)
			todos.GET("/:id", controllers.Todo.Show.Handle)
			todos.POST("", controllers.Todo.Store.Handle)
			todos.PUT("/:id", controllers.Todo.Update.Handle)
			todos.DELETE("/:id", controllers.Todo.Delete.Handle)
		}

		tags := api.Group("/tags")
		{
			tags.GET("", controllers.Tag.List.Handle)
			tags.GET("/:id", controllers.Tag.Show.Handle)
			tags.POST("", controllers.Tag.Store.Handle)
			tags.PUT("/:id", controllers.Tag.Update.Handle)
			tags.DELETE("/:id", controllers.Tag.Delete.Handle)
		}
	}
}
