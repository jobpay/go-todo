package todo

import (
	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type ShowUseCase struct {
	todoRepo repository.TodoRepository
}

func NewShowUseCase(todoRepo repository.TodoRepository) *ShowUseCase {
	return &ShowUseCase{
		todoRepo: todoRepo,
	}
}

func (u *ShowUseCase) Execute(id int) (*todo.Todo, error) {
	todoID, err := valueobject.NewID(id)
	if err != nil {
		return nil, err
	}
	return u.todoRepo.FindByID(todoID)
}
