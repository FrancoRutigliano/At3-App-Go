package infraSqlxRepository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (i *ImplSqlx) ResetPassword(email, newPassword string, db *sqlx.DB) (bool, error) {
	query := `UPDATE "public".users SET password = $1 WHERE email = $2`

	result, err := db.Exec(query, newPassword, email)
	if err != nil {
		return false, err
	}

	rowsA, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to check rows affected: %v", err)
	}

	if rowsA == 0 {
		return false, err
	}

	return true, nil
}
