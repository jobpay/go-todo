package todo

import (
	"net/http"

	"github.com/jobpay/todo/internal/application/usecase/todo"
	todoRequest "github.com/jobpay/todo/internal/presentation/request/todo"
	todoResponse "github.com/jobpay/todo/internal/presentation/response/todo"
	"github.com/labstack/echo/v4"
)

type StoreController struct {
	storeUseCase *todo.StoreUseCase
}

func NewStoreController(storeUseCase *todo.StoreUseCase) *StoreController {
	return &StoreController{
		storeUseCase: storeUseCase,
	}
}

func (c *StoreController) Handle(ctx echo.Context) error {
	var req todoRequest.StoreRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if err := req.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	input := todo.StoreInput{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
	}
	todoEntity, err := c.storeUseCase.Execute(input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	response := todoResponse.FromEntity(todoEntity)
	return ctx.JSON(http.StatusCreated, response)
}
