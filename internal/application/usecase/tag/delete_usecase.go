package tag

import (
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type DeleteUseCase struct {
	tagRepo repository.TagRepository
}

func NewDeleteUseCase(tagRepo repository.TagRepository) *DeleteUseCase {
	return &DeleteUseCase{tagRepo: tagRepo}
}

func (u *DeleteUseCase) Execute(id int) error {
	tagID, err := valueobject.NewID(id)
	if err != nil {
		return err
	}
	return u.tagRepo.Delete(tagID)
}
