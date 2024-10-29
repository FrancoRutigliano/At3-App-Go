package infraSqlxRepository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (i *ImplSqlx) FindByIdUpdate(id string, db *sqlx.DB) (bool, error) {
	query := `UPDATE "public".users SET role = 2 WHERE id = $1`

	result, err := db.Exec(query, id)
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
