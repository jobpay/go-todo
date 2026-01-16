package auth

import (
	"errors"

	"github.com/jobpay/todo/internal/domain/entity/user"
	"github.com/jobpay/todo/internal/domain/entity/user/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
	"gorm.io/gorm"
)

type RegisterUseCase struct {
	userRepo repository.UserRepository
}

type RegisterInput struct {
	Email string
	Name  string
}

func NewRegisterUseCase(userRepo repository.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{userRepo: userRepo}
}

func (u *RegisterUseCase) Execute(input RegisterInput) (*user.User, error) {
	email, err := valueobject.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	name, err := valueobject.NewName(input.Name)
	if err != nil {
		return nil, err
	}

	// メールアドレスの重複チェック
	existingUser, err := u.userRepo.FindByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	newUser := user.NewUser(email, name)

	if err := u.userRepo.Save(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}
