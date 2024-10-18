package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	Postgre_User     string
	Postgre_Password string
	Postgre_Db       string
	Postgre_Host     string
	Route_Migrate    string
}

func SetUp() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("port is empty")
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		return nil, errors.New("userDb is empty")
	}

	pass := os.Getenv("POSTGRES_PASSWORD")
	if pass == "" {
		return nil, errors.New("passDb is empty")
	}

	db := os.Getenv("POSTGRES_DB")
	if db == "" {
		return nil, errors.New("db is empty")
	}

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		return nil, errors.New("host is empty")
	}

	migrate := os.Getenv("ROUTE_MIGRATE")
	if migrate == "" {
		return nil, errors.New("route migrate is empty")
	}

	return &Config{
		Port:             port,
		Postgre_User:     user,
		Postgre_Password: pass,
		Postgre_Db:       db,
		Postgre_Host:     host,
		Route_Migrate:    migrate,
	}, nil
}
