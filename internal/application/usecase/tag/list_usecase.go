package tag

import (
	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/repository"
)

type ListUseCase struct {
	repo repository.TagRepository
}

func NewListUseCase(repo repository.TagRepository) *ListUseCase {
	return &ListUseCase{repo: repo}
}

func (uc *ListUseCase) Execute() ([]*tag.Tag, error) {
	return uc.repo.FindAll()
}

