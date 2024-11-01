package infraSqlxRepository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (i *ImplSqlx) FindByEmail(emailParam string, db *sqlx.DB) (bool, error) {
	query := `SELECT u.email FROM "public".users AS u WHERE u.email = $1 LIMIT 1;`

	var email string
	if err := db.Get(&email, query, emailParam); err != nil {
		// Si no se encuentran filas, significa que el correo no existe
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("error executing query: %w", err)
	}

	return true, nil
}
