package todo

import (
	"errors"
	"time"

	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type StoreUseCase struct {
	todoRepo repository.TodoRepository
}

type StoreInput struct {
	Title       string
	Description string
	DueDate     time.Time
}

func NewStoreUseCase(todoRepo repository.TodoRepository) *StoreUseCase {
	return &StoreUseCase{
		todoRepo: todoRepo,
	}
}

func (u *StoreUseCase) Execute(input StoreInput) (*todo.Todo, error) {
	title, err := valueobject.NewTitle(input.Title)
	if err != nil {
		return nil, err
	}

	description, err := valueobject.NewDescription(input.Description)
	if err != nil {
		return nil, err
	}

	if input.DueDate.Before(time.Now()) {
		return nil, errors.New("due date must be in the future")
	}

	newTodo := todo.NewTodo(title, description, input.DueDate)

	if err := u.todoRepo.Save(newTodo); err != nil {
		return nil, err
	}

	return newTodo, nil
}
