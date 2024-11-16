package infraSqlxRepository

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

func (i *ImplSqlx) GetUser(paylod authDto.LoginRequest, db *sqlx.DB) (authDto.LoginResponse, error) {
	query := `SELECT u.id, u.email, u.password, u.role FROM "public".users AS u
	WHERE email = $1 LIMIT 1;`

	var user authDto.LoginResponse

	if err := db.Get(&user, query, paylod.Email); err != nil {
		if err == sql.ErrNoRows {
			return authDto.LoginResponse{}, errors.New("not_found")
		}
		return authDto.LoginResponse{}, err
	}

	return user, nil
}
