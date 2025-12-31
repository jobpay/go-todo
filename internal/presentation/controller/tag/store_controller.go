package tag

import (
	"net/http"

	"github.com/labstack/echo/v4"

	tagUseCase "github.com/jobpay/todo/internal/application/usecase/tag"
	"github.com/jobpay/todo/internal/presentation/request"
	tagRequest "github.com/jobpay/todo/internal/presentation/request/tag"
	tagResponse "github.com/jobpay/todo/internal/presentation/response/tag"
)

type StoreController struct {
	useCase *tagUseCase.StoreUseCase
}

func NewStoreController(useCase *tagUseCase.StoreUseCase) *StoreController {
	return &StoreController{useCase: useCase}
}

func (ctrl *StoreController) Handle(c echo.Context) error {
	var req tagRequest.StoreRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if err := req.Validate(); err != nil {
		errors := request.ParseValidationErrors(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
	}

	result, err := ctrl.useCase.Execute(req.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, tagResponse.FromEntity(result))
}

