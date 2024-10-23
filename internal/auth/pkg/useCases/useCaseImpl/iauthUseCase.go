package authUseCaseImpl

import (
	infraSqlxRepository "at3-back/internal/auth/infrastructure/repository"
	authDto "at3-back/internal/auth/pkg/domain/dto"
	httpresponse "at3-back/pkg/httpResponse"

	"github.com/jmoiron/sqlx"
)

type IauthUseCase interface {
	Register(authDto.RegisterUser) httpresponse.ApiResponse
}

type Auth struct {
	Repository infraSqlxRepository.SqlxRepository
	Db         *sqlx.DB
}
