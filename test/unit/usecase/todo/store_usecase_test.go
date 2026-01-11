package todo_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jobpay/todo/internal/application/usecase/todo"
	domainTag "github.com/jobpay/todo/internal/domain/entity/tag"
	tagValueObject "github.com/jobpay/todo/internal/domain/entity/tag/valueobject"
	"github.com/jobpay/todo/test/mock"
)

func TestStoreUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoRepo := mock.NewMockTodoRepository(ctrl)
	mockTagRepo := mock.NewMockTagRepository(ctrl)
	useCase := todo.NewStoreUseCase(mockTodoRepo, mockTagRepo)

	t.Run("正常系: TODOを作成できる", func(t *testing.T) {
		input := todo.StoreInput{
			Title:       "Test TODO",
			Description: "Test Description",
			DueDate:     time.Now().Add(24 * time.Hour),
		}

		mockTodoRepo.EXPECT().Save(gomock.Any()).Return(nil)

		result, err := useCase.Execute(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.Title.String() != input.Title {
			t.Errorf("Expected title %s, got %s", input.Title, result.Title.String())
		}
	})

	t.Run("正常系: タグ付きTODOを作成できる", func(t *testing.T) {
		input := todo.StoreInput{
			Title:       "Test TODO with Tags",
			Description: "Test Description",
			DueDate:     time.Now().Add(24 * time.Hour),
			TagIDs:      []int{1, 2},
		}

		tagID1, _ := tagValueObject.NewID(1)
		tagTitle1, _ := tagValueObject.NewTitle("tag1")
		tag1 := &domainTag.Tag{
			ID:    tagID1,
			Title: tagTitle1,
		}

		tagID2, _ := tagValueObject.NewID(2)
		tagTitle2, _ := tagValueObject.NewTitle("tag2")
		tag2 := &domainTag.Tag{
			ID:    tagID2,
			Title: tagTitle2,
		}

		mockTagRepo.EXPECT().FindByID(tagID1).Return(tag1, nil)
		mockTagRepo.EXPECT().FindByID(tagID2).Return(tag2, nil)
		mockTodoRepo.EXPECT().Save(gomock.Any()).Return(nil)

		result, err := useCase.Execute(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(result.Tags) != 2 {
			t.Errorf("Expected 2 tags, got %d", len(result.Tags))
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

		mockTodoRepo.EXPECT().Save(gomock.Any()).Return(errors.New("database error"))

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: 存在しないタグIDでエラー", func(t *testing.T) {
		input := todo.StoreInput{
			Title:       "Test TODO",
			Description: "Test Description",
			DueDate:     time.Now().Add(24 * time.Hour),
			TagIDs:      []int{999},
		}

		tagID, _ := tagValueObject.NewID(999)
		mockTagRepo.EXPECT().FindByID(tagID).Return(nil, errors.New("tag not found"))

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
