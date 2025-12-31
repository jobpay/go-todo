package tag

import (
	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type StoreUseCase struct {
	repo repository.TagRepository
}

func NewStoreUseCase(repo repository.TagRepository) *StoreUseCase {
	return &StoreUseCase{repo: repo}
}

func (uc *StoreUseCase) Execute(title string) (*tag.Tag, error) {
	titleVO, err := valueobject.NewTitle(title)
	if err != nil {
		return nil, err
	}

	newTag := tag.NewTag(titleVO)

	if err := uc.repo.Save(newTag); err != nil {
		return nil, err
	}

	return newTag, nil
}

