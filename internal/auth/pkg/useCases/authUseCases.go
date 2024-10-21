package authUseCases

import authUseCaseImpl "at3-back/internal/auth/pkg/useCases/useCaseImpl"

type AuthImpl struct {
	impl authUseCaseImpl.IauthUseCase
}

func (a *AuthImpl) New() {
	// se inicializa repositorio
	// new repositorio
	//instancia base de datos

	// llamar estructura que implementa
}
