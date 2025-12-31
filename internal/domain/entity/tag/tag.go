package tag

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
)

type Tag struct {
	ID        valueobject.ID
	Title     valueobject.Title
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTag(title valueobject.Title) *Tag {
	now := time.Now()
	return &Tag{
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (t *Tag) Update(title valueobject.Title) {
	t.Title = title
	t.UpdatedAt = time.Now()
}

