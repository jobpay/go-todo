package tag

import (
	"github.com/jobpay/todo/internal/presentation/request"
)

type StoreRequest struct {
	Title string `json:"title" validate:"required,min=1,max=100"`
}

func (r *StoreRequest) Validate() error {
	return request.Validate(r)
}
