package tag

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
)

type TagModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (TagModel) TableName() string {
	return "tags"
}

func (m *TagModel) ToEntity() *tag.Tag {
	return &tag.Tag{
		ID:        valueobject.ID(m.ID),
		Title:     valueobject.Title(m.Title),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func FromEntity(entity *tag.Tag) *TagModel {
	return &TagModel{
		ID:        entity.ID.Int(),
		Title:     entity.Title.String(),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

