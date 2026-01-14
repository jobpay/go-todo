package user

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/user"
	"github.com/jobpay/todo/internal/domain/entity/user/valueobject"
)

type UserModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	Name      string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

func (UserModel) TableName() string {
	return "users"
}

func (m *UserModel) ToEntity() *user.User {
	id, _ := valueobject.NewID(m.ID)
	email, _ := valueobject.NewEmail(m.Email)
	name, _ := valueobject.NewName(m.Name)

	return &user.User{
		ID:        id,
		Email:     email,
		Name:      name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func FromEntity(e *user.User) *UserModel {
	return &UserModel{
		ID:        e.ID.Int(),
		Email:     e.Email.String(),
		Name:      e.Name.String(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

