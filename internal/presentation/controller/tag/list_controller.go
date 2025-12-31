package tag

import (
	"net/http"

	"github.com/labstack/echo/v4"

	tagUseCase "github.com/jobpay/todo/internal/application/usecase/tag"
	tagResponse "github.com/jobpay/todo/internal/presentation/response/tag"
)

type ListController struct {
	useCase *tagUseCase.ListUseCase
}

func NewListController(useCase *tagUseCase.ListUseCase) *ListController {
	return &ListController{useCase: useCase}
}

func (ctrl *ListController) Handle(c echo.Context) error {
	results, err := ctrl.useCase.Execute()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	responses := make([]*tagResponse.TagResponse, len(results))
	for i, result := range results {
		responses[i] = tagResponse.FromEntity(result)
	}

	return c.JSON(http.StatusOK, responses)
}

