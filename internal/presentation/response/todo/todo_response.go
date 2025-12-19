package todo

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/todo"
)

type TodoResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromEntity(entity *todo.Todo) *TodoResponse {
	return &TodoResponse{
		ID:          entity.ID.Int(),
		Title:       entity.Title.String(),
		Description: entity.Description.String(),
		Completed:   entity.Status.Bool(),
		DueDate:     entity.DueDate,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func FromEntities(entities []*todo.Todo) []*TodoResponse {
	responses := make([]*TodoResponse, len(entities))
	for i, entity := range entities {
		responses[i] = FromEntity(entity)
	}
	return responses
}
