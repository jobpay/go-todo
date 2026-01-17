package auth

import (
	"net/http"
	"strconv"

	"github.com/jobpay/todo/internal/application/usecase/auth"
	authResponse "github.com/jobpay/todo/internal/presentation/response/auth"
	"github.com/labstack/echo/v4"
)

type GetMeController struct {
	getMeUseCase *auth.GetMeUseCase
}

func NewGetMeController(getMeUseCase *auth.GetMeUseCase) *GetMeController {
	return &GetMeController{
		getMeUseCase: getMeUseCase,
	}
}

func (c *GetMeController) Handle(ctx echo.Context) error {
	// TODO: Phase 3でJWTから user_id を取得する
	// 現在は一時的にクエリパラメータで受け取る（テスト用）
	userIDStr := ctx.QueryParam("user_id")
	if userIDStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "user_id query parameter is required (temporary until JWT is implemented)",
		})
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid user_id",
		})
	}

	userEntity, err := c.getMeUseCase.Execute(userID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	response := authResponse.FromEntity(userEntity)
	return ctx.JSON(http.StatusOK, response)
}
