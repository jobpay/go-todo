package tag

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	tagUseCase "github.com/jobpay/todo/internal/application/usecase/tag"
)

type DeleteController struct {
	useCase *tagUseCase.DeleteUseCase
}

func NewDeleteController(useCase *tagUseCase.DeleteUseCase) *DeleteController {
	return &DeleteController{useCase: useCase}
}

func (ctrl *DeleteController) Handle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID",
		})
	}

	err = ctrl.useCase.Execute(id)
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

	return c.NoContent(http.StatusNoContent)
}

