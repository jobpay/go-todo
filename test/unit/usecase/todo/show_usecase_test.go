package todo_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jobpay/todo/internal/application/usecase/todo"
	todoEntity "github.com/jobpay/todo/internal/domain/entity/todo"
	"github.com/jobpay/todo/internal/domain/entity/todo/valueobject"
	"github.com/jobpay/todo/test/mock"
)

func TestShowUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTodoRepository(ctrl)
	useCase := todo.NewShowUseCase(mockRepo)

	t.Run("正常系: IDでTODOを取得できる", func(t *testing.T) {
		expectedTodo := &todoEntity.Todo{
			ID:          valueobject.ID(1),
			Title:       valueobject.Title("Test TODO"),
			Description: valueobject.Description("Test Description"),
			DueDate:     time.Now(),
		}

		mockRepo.EXPECT().FindByID(valueobject.ID(1)).Return(expectedTodo, nil)

		result, err := useCase.Execute(1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.ID != expectedTodo.ID {
			t.Errorf("Expected ID %d, got %d", expectedTodo.ID.Int(), result.ID.Int())
		}
	})

	t.Run("異常系: TODOが見つからない", func(t *testing.T) {
		mockRepo.EXPECT().FindByID(valueobject.ID(999)).Return(nil, errors.New("todo not found"))

		_, err := useCase.Execute(999)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
