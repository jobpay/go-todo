package tag

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	tagUseCase "github.com/jobpay/todo/internal/application/usecase/tag"
	"github.com/jobpay/todo/internal/presentation/request"
	tagRequest "github.com/jobpay/todo/internal/presentation/request/tag"
	tagResponse "github.com/jobpay/todo/internal/presentation/response/tag"
)

type UpdateController struct {
	useCase *tagUseCase.UpdateUseCase
}

func NewUpdateController(useCase *tagUseCase.UpdateUseCase) *UpdateController {
	return &UpdateController{useCase: useCase}
}

func (ctrl *UpdateController) Handle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID",
		})
	}

	var req tagRequest.UpdateRequest
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

	input := tagUseCase.UpdateInput{
		ID:    id,
		Title: req.Title,
	}

	result, err := ctrl.useCase.Execute(input)
	if err != nil {
		if err.Error() == "tag not found" {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, tagResponse.FromEntity(result))
}
