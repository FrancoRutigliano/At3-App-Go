package httpserver

import "github.com/gofiber/fiber/v2"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	app := fiber.New()
	return app.Listen(":8080")
}
