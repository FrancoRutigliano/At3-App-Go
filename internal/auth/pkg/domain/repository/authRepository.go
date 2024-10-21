package authRepository

type Repository interface {
	Login()
	Register()
	Reset()
}
