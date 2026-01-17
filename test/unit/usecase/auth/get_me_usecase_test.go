package auth_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jobpay/todo/internal/application/usecase/auth"
	userEntity "github.com/jobpay/todo/internal/domain/entity/user"
	userValueObject "github.com/jobpay/todo/internal/domain/entity/user/valueobject"
	"github.com/jobpay/todo/test/mock"
)

func TestGetMeUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepository(ctrl)
	useCase := auth.NewGetMeUseCase(mockRepo)

	t.Run("正常系: IDでユーザーを取得できる", func(t *testing.T) {
		userID, _ := userValueObject.NewID(1)
		email, _ := userValueObject.NewEmail("test@example.com")
		name, _ := userValueObject.NewName("Test User")

		expectedUser := &userEntity.User{
			ID:    userID,
			Email: email,
			Name:  name,
		}

		mockRepo.EXPECT().FindByID(userID).Return(expectedUser, nil)

		result, err := useCase.Execute(1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.ID.Int() != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID.Int())
		}
		if result.Email.String() != "test@example.com" {
			t.Errorf("Expected email test@example.com, got %s", result.Email.String())
		}
	})

	t.Run("異常系: ユーザーが見つからない", func(t *testing.T) {
		userID, _ := userValueObject.NewID(999)
		mockRepo.EXPECT().FindByID(userID).Return(nil, errors.New("user not found"))

		_, err := useCase.Execute(999)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
