package authUseCaseImpl

type IauthUseCase interface {
	Login()
	Register()
	Reset()
}

type Auth struct {
	// repositorio
	// instancia a db
}
