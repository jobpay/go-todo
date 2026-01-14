package user

import (
	"gorm.io/gorm"

	"github.com/jobpay/todo/internal/domain/entity/user"
	"github.com/jobpay/todo/internal/domain/entity/user/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(id valueobject.ID) (*user.User, error) {
	var model UserModel
	if err := r.db.First(&model, id.Int()).Error; err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

func (r *userRepository) FindByEmail(email valueobject.Email) (*user.User, error) {
	var model UserModel
	if err := r.db.Where("email = ?", email.String()).First(&model).Error; err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

func (r *userRepository) Save(user *user.User) error {
	model := FromEntity(user)
	if err := r.db.Create(model).Error; err != nil {
		return err
	}

	id, _ := valueobject.NewID(model.ID)
	user.ID = id
	return nil
}

func (r *userRepository) Update(user *user.User) error {
	model := FromEntity(user)
	return r.db.Save(model).Error
}

