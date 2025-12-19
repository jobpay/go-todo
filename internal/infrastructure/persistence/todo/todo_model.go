package todo

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
)

type TodoModel struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Completed   bool      `gorm:"not null;default:false"`
	DueDate     time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (TodoModel) TableName() string {
	return "todos"
}

func (m *TodoModel) ToEntity() *todo.Todo {
	return &todo.Todo{
		ID:          valueobject.ID(m.ID),
		Title:       valueobject.Title(m.Title),
		Description: valueobject.Description(m.Description),
		Status:      valueobject.FromBool(m.Completed),
		DueDate:     m.DueDate,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func FromEntity(entity *todo.Todo) *TodoModel {
	return &TodoModel{
		ID:          entity.ID.Int(),
		Title:       entity.Title.String(),
		Description: entity.Description.String(),
		Completed:   entity.Status.Bool(),
		DueDate:     entity.DueDate,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
