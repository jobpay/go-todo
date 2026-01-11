package todo

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/todo"
)

type Tag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type TodoResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	DueDate     time.Time `json:"due_date"`
	Tags        []*Tag    `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromEntity(entity *todo.Todo) *TodoResponse {
	tags := make([]*Tag, len(entity.Tags))
	for i, tag := range entity.Tags {
		tags[i] = &Tag{
			ID:    tag.ID.Int(),
			Title: tag.Title.String(),
		}
	}

	return &TodoResponse{
		ID:          entity.ID.Int(),
		Title:       entity.Title.String(),
		Description: entity.Description.String(),
		Completed:   entity.Status.Bool(),
		DueDate:     entity.DueDate,
		Tags:        tags,
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
