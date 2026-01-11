package todo

import (
	"time"

	"github.com/jobpay/todo/internal/domain/entity/tag"
	tagValueObject "github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type UpdateUseCase struct {
	todoRepo repository.TodoRepository
	tagRepo  repository.TagRepository
}

type UpdateInput struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	DueDate     time.Time
	TagIDs      []int
}

func NewUpdateUseCase(todoRepo repository.TodoRepository, tagRepo repository.TagRepository) *UpdateUseCase {
	return &UpdateUseCase{
		todoRepo: todoRepo,
		tagRepo:  tagRepo,
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

	existingTodo.Update(title, description, input.Completed, input.DueDate, tags)

	if err := u.todoRepo.Update(existingTodo); err != nil {
		return nil, err
	}

	return existingTodo, nil
}
