package data

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection() (*sqlx.DB, error) {
	connstr := os.Getenv("CONNECTION")
	db, err := sqlx.Connect("postgres", connstr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
