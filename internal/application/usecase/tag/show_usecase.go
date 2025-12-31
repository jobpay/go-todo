package tag

import (
	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type ShowUseCase struct {
	repo repository.TagRepository
}

func NewShowUseCase(repo repository.TagRepository) *ShowUseCase {
	return &ShowUseCase{repo: repo}
}

func (uc *ShowUseCase) Execute(id int) (*tag.Tag, error) {
	tagID, err := valueobject.NewID(id)
	if err != nil {
		return nil, err
	}

	tag, err := uc.repo.FindByID(tagID)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

