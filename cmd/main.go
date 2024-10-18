package main

import (
	httpserver "at3-back/internal/shared/infrastructure/entrypoint/http_server"
	"log"
)

func main() {
	s := httpserver.NewServer()
	err := s.Run()
	if err != nil {
		log.Fatal("error starting server")
	}

	log.Println("starting server...")
}
