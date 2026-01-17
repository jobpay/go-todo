package auth

import "github.com/jobpay/todo/internal/presentation/request"

type RegisterRequest struct {
	Email string `json:"email" validate:"required,email,max=255"`
	Name  string `json:"name" validate:"required,min=1,max=100"`
}

func (r *RegisterRequest) Validate() error {
	return request.Validate(r)
}
