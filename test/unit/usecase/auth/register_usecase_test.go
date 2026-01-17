package auth_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jobpay/todo/internal/application/usecase/auth"
	domainUser "github.com/jobpay/todo/internal/domain/entity/user"
	userValueObject "github.com/jobpay/todo/internal/domain/entity/user/valueobject"
	"github.com/jobpay/todo/test/mock"
	"gorm.io/gorm"
)

func TestRegisterUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock.NewMockUserRepository(ctrl)
	useCase := auth.NewRegisterUseCase(mockUserRepo)

	t.Run("正常系: ユーザーを登録できる", func(t *testing.T) {
		input := auth.RegisterInput{
			Email: "test@example.com",
			Name:  "Test User",
		}

		email, _ := userValueObject.NewEmail(input.Email)
		mockUserRepo.EXPECT().FindByEmail(email).Return(nil, gorm.ErrRecordNotFound)
		mockUserRepo.EXPECT().Save(gomock.Any()).Return(nil)

		result, err := useCase.Execute(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.Email.String() != input.Email {
			t.Errorf("Expected email %s, got %s", input.Email, result.Email.String())
		}
		if result.Name.String() != input.Name {
			t.Errorf("Expected name %s, got %s", input.Name, result.Name.String())
		}
	})

	t.Run("異常系: メールアドレスが空の場合エラー", func(t *testing.T) {
		input := auth.RegisterInput{
			Email: "",
			Name:  "Test User",
		}

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: メールアドレスが不正な形式", func(t *testing.T) {
		input := auth.RegisterInput{
			Email: "invalid-email",
			Name:  "Test User",
		}

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: 名前が空の場合エラー", func(t *testing.T) {
		input := auth.RegisterInput{
			Email: "test@example.com",
			Name:  "",
		}

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: 名前が100文字を超える場合エラー", func(t *testing.T) {
		longName := ""
		for i := 0; i < 101; i++ {
			longName += "a"
		}

		input := auth.RegisterInput{
			Email: "test@example.com",
			Name:  longName,
		}

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: 既に登録済みのメールアドレス", func(t *testing.T) {
		input := auth.RegisterInput{
			Email: "existing@example.com",
			Name:  "Test User",
		}

		email, _ := userValueObject.NewEmail(input.Email)
		existingUser := &domainUser.User{
			Email: email,
		}

		mockUserRepo.EXPECT().FindByEmail(email).Return(existingUser, nil)

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
		if err.Error() != "email already exists" {
			t.Errorf("Expected 'email already exists' error, got %v", err)
		}
	})

	t.Run("異常系: FindByEmailでデータベースエラー", func(t *testing.T) {
		input := auth.RegisterInput{
			Email: "test@example.com",
			Name:  "Test User",
		}

		email, _ := userValueObject.NewEmail(input.Email)
		mockUserRepo.EXPECT().FindByEmail(email).Return(nil, errors.New("database connection error"))

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("異常系: Saveでリポジトリエラー", func(t *testing.T) {
		input := auth.RegisterInput{
			Email: "test@example.com",
			Name:  "Test User",
		}

		email, _ := userValueObject.NewEmail(input.Email)
		mockUserRepo.EXPECT().FindByEmail(email).Return(nil, gorm.ErrRecordNotFound)
		mockUserRepo.EXPECT().Save(gomock.Any()).Return(errors.New("database error"))

		_, err := useCase.Execute(input)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
