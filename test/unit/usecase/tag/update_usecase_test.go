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

func TestUpdateUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTagRepository(ctrl)
	useCase := tag.NewUpdateUseCase(mockRepo)

	t.Run("正常系: タグを更新できる", func(t *testing.T) {
		input := tag.UpdateInput{
			ID:    1,
			Title: "high-priority",
		}

		existingTag := &entityTag.Tag{
			ID:    valueobject.ID(1),
			Title: valueobject.Title("urgent"),
		}

		mockRepo.EXPECT().
			FindByID(valueobject.ID(1)).
			Return(existingTag, nil)

		mockRepo.EXPECT().Update(gomock.Any()).Return(nil)

		result, err := useCase.Execute(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.Title.String() != input.Title {
			t.Errorf("Expected title %s, got %s", input.Title, result.Title.String())
		}
	})

	t.Run("異常系: 無効なID", func(t *testing.T) {
		input := tag.UpdateInput{
			ID:    -1,
			Title: "test",
		}

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: タグが見つからない", func(t *testing.T) {
		input := tag.UpdateInput{
			ID:    999,
			Title: "test",
		}

		mockRepo.EXPECT().
			FindByID(valueobject.ID(999)).
			Return(nil, errors.New("tag not found"))

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: タイトルが空", func(t *testing.T) {
		input := tag.UpdateInput{
			ID:    1,
			Title: "",
		}

		existingTag := &entityTag.Tag{
			ID:    valueobject.ID(1),
			Title: valueobject.Title("urgent"),
		}

		mockRepo.EXPECT().
			FindByID(valueobject.ID(1)).
			Return(existingTag, nil)

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

