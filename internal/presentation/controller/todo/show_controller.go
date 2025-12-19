package todo

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/jobpay/todo/internal/application/usecase/todo"
	todoResponse "github.com/jobpay/todo/internal/presentation/response/todo"
)

type ShowController struct {
	showUseCase *todo.ShowUseCase
}

func NewShowController(showUseCase *todo.ShowUseCase) *ShowController {
	return &ShowController{
		showUseCase: showUseCase,
	}
}

func (c *ShowController) Handle(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	todoEntity, err := c.showUseCase.Execute(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	response := todoResponse.FromEntity(todoEntity)
	return ctx.JSON(http.StatusOK, response)
}
