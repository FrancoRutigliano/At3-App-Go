package data

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection() (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", os.Getenv("CONNECTION"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
