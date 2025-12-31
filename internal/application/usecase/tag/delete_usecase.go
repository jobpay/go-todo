package tag

import (
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type DeleteUseCase struct {
	repo repository.TagRepository
}

func NewDeleteUseCase(repo repository.TagRepository) *DeleteUseCase {
	return &DeleteUseCase{repo: repo}
}

func (uc *DeleteUseCase) Execute(id int) error {
	tagID, err := valueobject.NewID(id)
	if err != nil {
		return err
	}

	return uc.repo.Delete(tagID)
}

