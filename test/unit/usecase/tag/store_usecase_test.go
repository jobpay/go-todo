package tag_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jobpay/todo/internal/application/usecase/tag"
	"github.com/jobpay/todo/test/mock"
)

func TestStoreUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTagRepository(ctrl)
	useCase := tag.NewStoreUseCase(mockRepo)

	t.Run("正常系: タグを作成できる", func(t *testing.T) {
		input := tag.StoreInput{
			Title: "urgent",
		}

		mockRepo.EXPECT().Save(gomock.Any()).Return(nil)

		result, err := useCase.Execute(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.Title.String() != input.Title {
			t.Errorf("Expected title %s, got %s", input.Title, result.Title.String())
		}
	})

	t.Run("異常系: タイトルが空の場合エラー", func(t *testing.T) {
		input := tag.StoreInput{
			Title: "",
		}

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: タイトルが100文字を超える場合エラー", func(t *testing.T) {
		input := tag.StoreInput{
			Title: "a" + string(make([]byte, 100)),
		}

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: リポジトリエラー", func(t *testing.T) {
		input := tag.StoreInput{
			Title: "urgent",
		}

		mockRepo.EXPECT().Save(gomock.Any()).Return(errors.New("database error"))

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
