package auth

import (
	"github.com/jobpay/todo/internal/domain/entity/user"
	"github.com/jobpay/todo/internal/domain/entity/user/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type GetMeUseCase struct {
	userRepo repository.UserRepository
}

func NewGetMeUseCase(userRepo repository.UserRepository) *GetMeUseCase {
	return &GetMeUseCase{userRepo: userRepo}
}

func (u *GetMeUseCase) Execute(userID int) (*user.User, error) {
	id, err := valueobject.NewID(userID)
	if err != nil {
		return nil, err
	}

	return u.userRepo.FindByID(id)
}
