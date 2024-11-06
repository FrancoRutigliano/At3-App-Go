package authRepository

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindByEmail(string, string, *sqlx.DB) (bool, error)
	CreateUserAccount(*authDto.RegisterDb, *sqlx.DB) error
	CreateCompanyAccount(*authDto.RegisterCompanyDB, *sqlx.DB) error
	FindByIdUpdate(string, *sqlx.DB) (bool, error)
	GetUser(authDto.LoginRequest, *sqlx.DB) (authDto.LoginResponse, error)
}
