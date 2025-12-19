package todo

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type UpdateUseCase struct {
	todoRepo repository.TodoRepository
}

type UpdateInput struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	DueDate     time.Time
}

func NewUpdateUseCase(todoRepo repository.TodoRepository) *UpdateUseCase {
	return &UpdateUseCase{
		todoRepo: todoRepo,
	}
}

func (u *UpdateUseCase) Execute(input UpdateInput) (*todo.Todo, error) {
	id, err := valueobject.NewID(input.ID)
	if err != nil {
		return nil, err
	}

	title, err := valueobject.NewTitle(input.Title)
	if err != nil {
		return nil, err
	}

	description, err := valueobject.NewDescription(input.Description)
	if err != nil {
		return nil, err
	}

	existingTodo, err := u.todoRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	existingTodo.Update(title, description, input.DueDate)

	if input.Completed {
		existingTodo.Complete()
	} else {
		existingTodo.Reopen()
	}

	if err := u.todoRepo.Update(existingTodo); err != nil {
		return nil, err
	}

	return existingTodo, nil
}
