package tag

import (
	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/repository"
)

type ListUseCase struct {
	tagRepo repository.TagRepository
}

func NewListUseCase(tagRepo repository.TagRepository) *ListUseCase {
	return &ListUseCase{tagRepo: tagRepo}
}

func (u *ListUseCase) Execute() ([]*tag.Tag, error) {
	return u.tagRepo.FindAll()
}
