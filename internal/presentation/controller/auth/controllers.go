package auth

type Controllers struct {
	Register *RegisterController
	GetMe    *GetMeController
}

func NewControllers(
	register *RegisterController,
	getMe *GetMeController,
) *Controllers {
	return &Controllers{
		Register: register,
		GetMe:    getMe,
	}
}
