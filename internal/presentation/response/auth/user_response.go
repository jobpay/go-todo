package auth

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/user"
)

type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromEntity(entity *user.User) *UserResponse {
	return &UserResponse{
		ID:        entity.ID.Int(),
		Email:     entity.Email.String(),
		Name:      entity.Name.String(),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
