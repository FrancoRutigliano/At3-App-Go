package httpserver

import (
	"at3-back/config"
	"at3-back/internal/shared/infrastructure/entrypoint/http_server/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Run() error {
	app := fiber.New()

	routes.Init(app)

	return app.Listen(os.Getenv("PORT"))
}
