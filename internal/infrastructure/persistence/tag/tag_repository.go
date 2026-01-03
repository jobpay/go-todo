package tag

import (
	"errors"

	"github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	impl "github.com/jobpay/todo/internal/domain/repository"
	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) impl.TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) Save(tag *tag.Tag) error {
	model := FromEntity(tag)
	if err := r.db.Create(model).Error; err != nil {
		return err
	}
	tag.ID = valueobject.ID(model.ID)
	return nil
}

func (r *tagRepository) FindByID(id valueobject.ID) (*tag.Tag, error) {
	var model TagModel
	if err := r.db.First(&model, id.Int()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tag not found")
		}
		return nil, err
	}
	return model.ToEntity(), nil
}

func (r *tagRepository) FindAll() ([]*tag.Tag, error) {
	var models []TagModel
	if err := r.db.Order("created_at DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	tags := make([]*tag.Tag, len(models))
	for i, model := range models {
		tags[i] = model.ToEntity()
	}
	return tags, nil
}

func (r *tagRepository) Update(tag *tag.Tag) error {
	model := FromEntity(tag)
	result := r.db.Model(&TagModel{}).Where("id = ?", tag.ID.Int()).Updates(model)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("tag not found")
	}
	return nil
}

func (r *tagRepository) Delete(id valueobject.ID) error {
	result := r.db.Delete(&TagModel{}, id.Int())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("tag not found")
	}
	return nil
}

