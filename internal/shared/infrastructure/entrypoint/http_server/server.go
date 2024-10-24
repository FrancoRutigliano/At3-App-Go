package httpserver

import (
	"at3-back/internal/shared/infrastructure/entrypoint/http_server/routes"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	app := fiber.New()

	routes.Init(app)

	return app.Listen(":8080")
}
