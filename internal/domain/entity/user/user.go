package user

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/user/valueobject"
)

type User struct {
	ID        valueobject.ID
	Email     valueobject.Email
	Name      valueobject.Name
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(email valueobject.Email, name valueobject.Name) *User {
	now := time.Now()
	return &User{
		Email:     email,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (u *User) UpdateName(name valueobject.Name) {
	u.Name = name
	u.UpdatedAt = time.Now()
}

