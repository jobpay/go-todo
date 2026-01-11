package todo

import (
	"errors"
	"time"

	"github.com/jobpay/todo/internal/domain/entity/tag"
	tagValueObject "github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type StoreUseCase struct {
	todoRepo repository.TodoRepository
	tagRepo  repository.TagRepository
}

type StoreInput struct {
	Title       string
	Description string
	DueDate     time.Time
	TagIDs      []int
}

func NewStoreUseCase(todoRepo repository.TodoRepository, tagRepo repository.TagRepository) *StoreUseCase {
	return &StoreUseCase{
		todoRepo: todoRepo,
		tagRepo:  tagRepo,
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

	// Fetch tags by IDs
	tags := make([]*tag.Tag, 0, len(input.TagIDs))
	for _, tagID := range input.TagIDs {
		tagIDVO, err := tagValueObject.NewID(tagID)
		if err != nil {
			return nil, err
		}
		tag, err := u.tagRepo.FindByID(tagIDVO)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	newTodo := todo.NewTodo(title, description, input.DueDate, tags)

	if err := u.todoRepo.Save(newTodo); err != nil {
		return nil, err
	}

	return newTodo, nil
}
