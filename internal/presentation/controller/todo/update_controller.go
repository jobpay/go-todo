package todo

import (
	"net/http"
	"strconv"

	"github.com/jobpay/todo/internal/application/usecase/todo"
	"github.com/jobpay/todo/internal/presentation/request"
	todoRequest "github.com/jobpay/todo/internal/presentation/request/todo"
	todoResponse "github.com/jobpay/todo/internal/presentation/response/todo"
	"github.com/labstack/echo/v4"
)

type UpdateController struct {
	updateUseCase *todo.UpdateUseCase
}

func NewUpdateController(updateUseCase *todo.UpdateUseCase) *UpdateController {
	return &UpdateController{
		updateUseCase: updateUseCase,
	}
}

func (c *UpdateController) Handle(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var req todoRequest.UpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if err := req.Validate(); err != nil {
		errors := request.ParseValidationErrors(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
	}

	input := todo.UpdateInput{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
		DueDate:     req.DueDate,
		TagIDs:      req.TagIDs,
	}
	todoEntity, err := c.updateUseCase.Execute(input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	response := todoResponse.FromEntity(todoEntity)
	return ctx.JSON(http.StatusOK, response)
}
