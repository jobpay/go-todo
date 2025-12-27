package todo

import (
	"errors"

	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
	impl "github.com/jobpay/todo/internal/domain/repository"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) impl.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Save(todo *todo.Todo) error {
	model := FromEntity(todo)
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	todo.ID = valueobject.ID(model.ID)
	return nil
}

func (r *todoRepository) FindByID(id valueobject.ID) (*todo.Todo, error) {
	var model TodoModel
	if err := r.db.First(&model, id.Int()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return model.ToEntity(), nil
}

func (r *todoRepository) FindAll() ([]*todo.Todo, error) {
	var models []TodoModel
	if err := r.db.Order("created_at DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	todos := make([]*todo.Todo, len(models))
	for i, model := range models {
		todos[i] = model.ToEntity()
	}
	return todos, nil
}

func (r *todoRepository) Update(todo *todo.Todo) error {
	model := FromEntity(todo)
	result := r.db.Model(&TodoModel{}).Where("id = ?", todo.ID.Int()).Updates(model)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}
	return nil
}

func (r *todoRepository) Delete(id valueobject.ID) error {
	result := r.db.Delete(&TodoModel{}, id.Int())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}
	return nil
}
