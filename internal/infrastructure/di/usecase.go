package di

import (
	authUseCase "github.com/jobpay/todo/internal/application/usecase/auth"
	tagUseCase "github.com/jobpay/todo/internal/application/usecase/tag"
	todoUseCase "github.com/jobpay/todo/internal/application/usecase/todo"
)

func (c *Container) provideUseCases() error {
	useCases := []interface{}{
		todoUseCase.NewShowUseCase,
		todoUseCase.NewListUseCase,
		todoUseCase.NewStoreUseCase,
		todoUseCase.NewUpdateUseCase,
		todoUseCase.NewDeleteUseCase,
		tagUseCase.NewShowUseCase,
		tagUseCase.NewListUseCase,
		tagUseCase.NewStoreUseCase,
		tagUseCase.NewUpdateUseCase,
		tagUseCase.NewDeleteUseCase,
		authUseCase.NewRegisterUseCase,
		authUseCase.NewGetMeUseCase,
	}

	for _, useCase := range useCases {
		if err := c.container.Provide(useCase); err != nil {
			return err
		}
	}

	return nil
}
