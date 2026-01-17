package repository

import (
	"github.com/jobpay/todo/internal/domain/entity/user"
	"github.com/jobpay/todo/internal/domain/entity/user/valueobject"
)

type UserRepository interface {
	FindByID(id valueobject.ID) (*user.User, error)
	FindByEmail(email valueobject.Email) (*user.User, error)
	Save(user *user.User) error
	Update(user *user.User) error
}

