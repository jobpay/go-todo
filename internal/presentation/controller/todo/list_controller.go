package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/jobpay/todo/internal/application/usecase/todo"
	todoResponse "github.com/jobpay/todo/internal/presentation/response/todo"
)

type ListController struct {
	listUseCase *todo.ListUseCase
}

func NewListController(listUseCase *todo.ListUseCase) *ListController {
	return &ListController{
		listUseCase: listUseCase,
	}
}

func (c *ListController) Handle(ctx echo.Context) error {
	todos, err := c.listUseCase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	response := todoResponse.FromEntities(todos)
	return ctx.JSON(http.StatusOK, response)
}
