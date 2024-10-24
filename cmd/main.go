package main

import (
	"at3-back/config"
	httpserver "at3-back/internal/shared/infrastructure/entrypoint/http_server"
	"log"
)

func main() {
	config, err := config.SetUp()
	if err != nil {
		log.Fatal("error loading .env --> ", err)
	}
	s := httpserver.NewServer(config)
	err = s.Run()
	if err != nil {
		log.Fatal("error starting server")
	}

	log.Println("starting server...")
}
