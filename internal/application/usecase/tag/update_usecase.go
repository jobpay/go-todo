package tag

import (
	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type UpdateUseCase struct {
	tagRepo repository.TagRepository
}

type UpdateInput struct {
	ID    int
	Title string
}

func NewUpdateUseCase(tagRepo repository.TagRepository) *UpdateUseCase {
	return &UpdateUseCase{tagRepo: tagRepo}
}

func (u *UpdateUseCase) Execute(input UpdateInput) (*tag.Tag, error) {
	id, err := valueobject.NewID(input.ID)
	if err != nil {
		return nil, err
	}

	tag, err := u.tagRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	title, err := valueobject.NewTitle(input.Title)
	if err != nil {
		return nil, err
	}

	tag.Update(title)

	if err := u.tagRepo.Update(tag); err != nil {
		return nil, err
	}

	return tag, nil
}
