package todo

import (
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type DeleteUseCase struct {
	todoRepo repository.TodoRepository
}

func NewDeleteUseCase(todoRepo repository.TodoRepository) *DeleteUseCase {
	return &DeleteUseCase{
		todoRepo: todoRepo,
	}
}

func (u *DeleteUseCase) Execute(id int) error {
	todoID, err := valueobject.NewID(id)
	if err != nil {
		return err
	}
	return u.todoRepo.Delete(todoID)
}
