package authUseCases

import (
	infraSqlxRepository "at3-back/internal/auth/infrastructure/repository"
	authUseCaseImpl "at3-back/internal/auth/pkg/useCases/useCaseImpl"
	"at3-back/internal/shared/infrastructure/data"
	"log"
)

type AuthImpl struct {
	Impl authUseCaseImpl.IauthUseCase
}

func (a *AuthImpl) New() {
	var repository infraSqlxRepository.SqlxRepository

	repository.New()

	db, err := data.GetConnection()
	if err != nil {
		log.Fatal("error to instance db")
	}

	a.Impl = &authUseCaseImpl.Auth{
		Repository: repository,
		Db:         db,
	}
}
