package tag

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/tag"
)

type TagResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromEntity(entity *tag.Tag) *TagResponse {
	return &TagResponse{
		ID:        entity.ID.Int(),
		Title:     entity.Title.String(),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

