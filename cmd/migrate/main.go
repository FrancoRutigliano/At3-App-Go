package main

import (
	"at3-back/config"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.SetUp()
	if err != nil {
		log.Fatal("err to load env variables --> ", err)
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", config.Postgre_User, config.Postgre_Password, config.Postgre_Host, config.Postgre_Db)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error to open db: ", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("error to inicialize driver: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(config.Route_Migrate, config.Postgre_Db, driver)
	if err != nil {
		log.Fatal("error creating migrate intance: ", err)
	}

	cmd := os.Args[len(os.Args)-1]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("error up migrations: ", err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("error down migrations")
		}
	default:
		log.Println("command not found")
	}

	log.Println("migration completed")
}
