package todo

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/jobpay/todo/internal/application/usecase/todo"
)

type DeleteController struct {
	deleteUseCase *todo.DeleteUseCase
}

func NewDeleteController(deleteUseCase *todo.DeleteUseCase) *DeleteController {
	return &DeleteController{
		deleteUseCase: deleteUseCase,
	}
}

func (c *DeleteController) Handle(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	if err := c.deleteUseCase.Execute(id); err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}
