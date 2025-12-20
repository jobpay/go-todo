package todo

import (
	"time"

	"github.com/jobpay/todo/internal/presentation/request"
)

type StoreRequest struct {
	Title       string    `json:"title" validate:"required,min=1,max=100"`
	Description string    `json:"description" validate:"max=500"`
	DueDate     time.Time `json:"due_date" validate:"required"`
}

func (r *StoreRequest) Validate() error {
	return request.Validate(r)
}
