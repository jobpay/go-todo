package todo

import (
	"time"

	"github.com/jobpay/todo/internal/presentation/request"
)

type UpdateRequest struct {
	Title       string    `json:"title" validate:"required,min=1,max=100"`
	Description string    `json:"description" validate:"max=500"`
	Completed   bool      `json:"completed"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	TagIDs      []int     `json:"tag_ids"`
}

func (r *UpdateRequest) Validate() error {
	return request.Validate(r)
}
