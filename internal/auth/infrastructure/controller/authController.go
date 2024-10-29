package authController

import authUseCases "at3-back/internal/auth/pkg/useCases"

type Auth struct {
	handler authUseCases.AuthImpl
}

func (a *Auth) New() {
	a.handler.New()
}
