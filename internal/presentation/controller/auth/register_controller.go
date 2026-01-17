package auth

import (
	"net/http"

	"github.com/jobpay/todo/internal/application/usecase/auth"
	"github.com/jobpay/todo/internal/presentation/request"
	authRequest "github.com/jobpay/todo/internal/presentation/request/auth"
	authResponse "github.com/jobpay/todo/internal/presentation/response/auth"
	"github.com/labstack/echo/v4"
)

type RegisterController struct {
	registerUseCase *auth.RegisterUseCase
}

func NewRegisterController(registerUseCase *auth.RegisterUseCase) *RegisterController {
	return &RegisterController{
		registerUseCase: registerUseCase,
	}
}

func (c *RegisterController) Handle(ctx echo.Context) error {
	var req authRequest.RegisterRequest
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

	input := auth.RegisterInput{
		Email: req.Email,
		Name:  req.Name,
	}
	userEntity, err := c.registerUseCase.Execute(input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	response := authResponse.FromEntity(userEntity)
	return ctx.JSON(http.StatusCreated, response)
}
