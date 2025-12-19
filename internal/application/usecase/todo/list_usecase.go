package todo

import (
	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/repository"
)

type ListUseCase struct {
	todoRepo repository.TodoRepository
}

func NewListUseCase(todoRepo repository.TodoRepository) *ListUseCase {
	return &ListUseCase{
		todoRepo: todoRepo,
	}
}

func (u *ListUseCase) Execute() ([]*todo.Todo, error) {
	return u.todoRepo.FindAll()
}
