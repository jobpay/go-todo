package todo

import "time"

type StoreRequest struct {
	Title       string    `json:"title" validate:"required,min=1,max=100"`
	Description string    `json:"description" validate:"max=500"`
	DueDate     time.Time `json:"due_date" validate:"required"`
}
