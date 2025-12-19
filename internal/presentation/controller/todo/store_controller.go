package todo

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/jobpay/todo/internal/application/usecase/todo"
	todoRequest "github.com/jobpay/todo/internal/presentation/request/todo"
	todoResponse "github.com/jobpay/todo/internal/presentation/response/todo"
)

type StoreController struct {
	storeUseCase *todo.StoreUseCase
	validator    *validator.Validate
}

func NewStoreController(storeUseCase *todo.StoreUseCase) *StoreController {
	return &StoreController{
		storeUseCase: storeUseCase,
		validator:    validator.New(),
	}
}

func (c *StoreController) Handle(ctx echo.Context) error {
	var req todoRequest.StoreRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if err := c.validator.Struct(req); err != nil {
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
