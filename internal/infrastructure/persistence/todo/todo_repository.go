package todo

import (
	"errors"

	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
	impl "github.com/jobpay/todo/internal/domain/repository"
	tagPersistence "github.com/jobpay/todo/internal/infrastructure/persistence/tag"
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

	tagModels := make([]tagPersistence.TagModel, len(todo.Tags))
	for i, tag := range todo.Tags {
		tagModels[i] = *tagPersistence.FromEntity(tag)
	}
	model.Tags = tagModels

	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	todo.ID = valueobject.ID(model.ID)
	return nil
}

func (r *todoRepository) FindByID(id valueobject.ID) (*todo.Todo, error) {
	var model TodoModel
	if err := r.db.Preload("Tags").First(&model, id.Int()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return model.ToEntity(), nil
}

func (r *todoRepository) FindAll() ([]*todo.Todo, error) {
	var models []TodoModel
	if err := r.db.Preload("Tags").Order("created_at DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	todos := make([]*todo.Todo, len(models))
	for i, model := range models {
		todos[i] = model.ToEntity()
	}
	return todos, nil
}

func (r *todoRepository) Update(todo *todo.Todo) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		model := FromEntity(todo)
		result := tx.Model(&TodoModel{}).Where("id = ?", todo.ID.Int()).Updates(model)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("todo not found")
		}

		if err := tx.Where("todo_id = ?", todo.ID.Int()).Delete(&TodoTagModel{}).Error; err != nil {
			return err
		}

		if len(todo.Tags) > 0 {
			tagModels := make([]TodoTagModel, len(todo.Tags))
			for i, tag := range todo.Tags {
				tagModels[i] = TodoTagModel{
					TodoID: todo.ID.Int(),
					TagID:  tag.ID.Int(),
				}
			}
			if err := tx.Create(&tagModels).Error; err != nil {
				return err
			}
		}

		return nil
	})
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
