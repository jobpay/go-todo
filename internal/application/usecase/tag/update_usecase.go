package tag

import (
	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type UpdateUseCase struct {
	repo repository.TagRepository
}

func NewUpdateUseCase(repo repository.TagRepository) *UpdateUseCase {
	return &UpdateUseCase{repo: repo}
}

func (uc *UpdateUseCase) Execute(id int, title string) (*tag.Tag, error) {
	tagID, err := valueobject.NewID(id)
	if err != nil {
		return nil, err
	}

	tag, err := uc.repo.FindByID(tagID)
	if err != nil {
		return nil, err
	}

	titleVO, err := valueobject.NewTitle(title)
	if err != nil {
		return nil, err
	}

	tag.Update(titleVO)

	if err := uc.repo.Update(tag); err != nil {
		return nil, err
	}

	return tag, nil
}

