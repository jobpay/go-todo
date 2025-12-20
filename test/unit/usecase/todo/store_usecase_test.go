package todo_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jobpay/todo/internal/application/usecase/todo"
	"github.com/jobpay/todo/test/mock"
)

func TestStoreUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTodoRepository(ctrl)
	useCase := todo.NewStoreUseCase(mockRepo)

	t.Run("正常系: TODOを作成できる", func(t *testing.T) {
		input := todo.StoreInput{
			Title:       "Test TODO",
			Description: "Test Description",
			DueDate:     time.Now().Add(24 * time.Hour),
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
		input := todo.StoreInput{
			Title:       "",
			Description: "Test Description",
			DueDate:     time.Now().Add(24 * time.Hour),
		}

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: 期限が過去の場合エラー", func(t *testing.T) {
		input := todo.StoreInput{
			Title:       "Test TODO",
			Description: "Test Description",
			DueDate:     time.Now().Add(-24 * time.Hour),
		}

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: リポジトリエラー", func(t *testing.T) {
		input := todo.StoreInput{
			Title:       "Test TODO",
			Description: "Test Description",
			DueDate:     time.Now().Add(24 * time.Hour),
		}

		mockRepo.EXPECT().Save(gomock.Any()).Return(errors.New("database error"))

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
