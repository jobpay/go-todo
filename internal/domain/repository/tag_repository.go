package repository

import (
	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
)

type TagRepository interface {
	FindByID(id valueobject.ID) (*tag.Tag, error)
	FindAll() ([]*tag.Tag, error)
	Save(tag *tag.Tag) error
	Update(tag *tag.Tag) error
	Delete(id valueobject.ID) error
}

