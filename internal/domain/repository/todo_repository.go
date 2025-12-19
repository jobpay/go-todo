package repository

import (
	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
)

type TodoRepository interface {
	Save(todo *todo.Todo) error
	FindByID(id valueobject.ID) (*todo.Todo, error)
	FindAll() ([]*todo.Todo, error)
	Update(todo *todo.Todo) error
	Delete(id valueobject.ID) error
}
