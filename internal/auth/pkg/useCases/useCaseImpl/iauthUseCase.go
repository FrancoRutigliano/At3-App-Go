package authUseCaseImpl

import (
	infraSqlxRepository "at3-back/internal/auth/infrastructure/repository"
	authDto "at3-back/internal/auth/pkg/domain/dto"
	"at3-back/internal/shared/infrastructure/service"
	httpresponse "at3-back/pkg/httpResponse"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type IauthUseCase interface {
	Register(authDto.RegisterUser) httpresponse.ApiResponse
	Confirm(string, string) httpresponse.ApiResponse
	Login(authDto.LoginRequest) httpresponse.ApiResponse
	RegisterCompany(authDto.RegisterCompanyRequest) httpresponse.ApiResponse
}

type Auth struct {
	Repository   infraSqlxRepository.SqlxRepository
	Db           *sqlx.DB
	Redis        *redis.Client
	Ctx          context.Context
	EmailService service.EmailService
	JwtService   service.JwtService
}
