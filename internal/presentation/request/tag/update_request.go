package tag

import (
	"github.com/jobpay/todo/internal/presentation/request"
)

type UpdateRequest struct {
	Title string `json:"title" validate:"required,min=1,max=100"`
}

func (r *UpdateRequest) Validate() error {
	return request.Validate(r)
}

