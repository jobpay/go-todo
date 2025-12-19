package todo

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
)

type Todo struct {
	ID          valueobject.ID
	Title       valueobject.Title
	Description valueobject.Description
	Status      valueobject.Status
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodo(title valueobject.Title, description valueobject.Description, dueDate time.Time) *Todo {
	now := time.Now()
	return &Todo{
		Title:       title,
		Description: description,
		Status:      valueobject.StatusPending,
		DueDate:     dueDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Todo) Complete() {
	t.Status = valueobject.StatusCompleted
	t.UpdatedAt = time.Now()
}

func (t *Todo) Reopen() {
	t.Status = valueobject.StatusPending
	t.UpdatedAt = time.Now()
}

func (t *Todo) Update(title valueobject.Title, description valueobject.Description, dueDate time.Time) {
	t.Title = title
	t.Description = description
	t.DueDate = dueDate
	t.UpdatedAt = time.Now()
}
