package tag

import (
	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/internal/domain/repository"
)

type ShowUseCase struct {
	tagRepo repository.TagRepository
}

func NewShowUseCase(tagRepo repository.TagRepository) *ShowUseCase {
	return &ShowUseCase{tagRepo: tagRepo}
}

func (u *ShowUseCase) Execute(id int) (*tag.Tag, error) {
	tagID, err := valueobject.NewID(id)
	if err != nil {
		return nil, err
	}
	return u.tagRepo.FindByID(tagID)
}
