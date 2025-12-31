package tag_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jobpay/todo/internal/application/usecase/tag"
	entityTag "github.com/jobpay/todo/internal/domain/entity/tag"
	"github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/test/mock"
)

func TestShowUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTagRepository(ctrl)
	useCase := tag.NewShowUseCase(mockRepo)

	t.Run("正常系: タグを取得できる", func(t *testing.T) {
		id := 1
		expectedTag := &entityTag.Tag{
			ID:    valueobject.ID(1),
			Title: valueobject.Title("urgent"),
		}

		mockRepo.EXPECT().
			FindByID(valueobject.ID(1)).
			Return(expectedTag, nil)

		result, err := useCase.Execute(id)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.ID != expectedTag.ID {
			t.Errorf("Expected ID %d, got %d", expectedTag.ID, result.ID)
		}
	})

	t.Run("異常系: 無効なID", func(t *testing.T) {
		_, err := useCase.Execute(-1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: タグが見つからない", func(t *testing.T) {
		id := 999
		mockRepo.EXPECT().
			FindByID(valueobject.ID(999)).
			Return(nil, errors.New("tag not found"))

		_, err := useCase.Execute(id)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

