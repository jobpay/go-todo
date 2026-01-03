package tag

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	tagUseCase "github.com/jobpay/todo/internal/application/usecase/tag"
	tagResponse "github.com/jobpay/todo/internal/presentation/response/tag"
)

type ShowController struct {
	useCase *tagUseCase.ShowUseCase
}

func NewShowController(useCase *tagUseCase.ShowUseCase) *ShowController {
	return &ShowController{useCase: useCase}
}

func (ctrl *ShowController) Handle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID",
		})
	}

	result, err := ctrl.useCase.Execute(id)
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

