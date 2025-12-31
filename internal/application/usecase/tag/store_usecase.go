package tag

import (
	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type StoreUseCase struct {
	tagRepo repository.TagRepository
}

type StoreInput struct {
	Title string
}

func NewStoreUseCase(tagRepo repository.TagRepository) *StoreUseCase {
	return &StoreUseCase{tagRepo: tagRepo}
}

func (u *StoreUseCase) Execute(input StoreInput) (*tag.Tag, error) {
	title, err := valueobject.NewTitle(input.Title)
	if err != nil {
		return nil, err
	}

	newTag := tag.NewTag(title)

	if err := u.tagRepo.Save(newTag); err != nil {
		return nil, err
	}

	return newTag, nil
}
