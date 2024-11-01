package httpserver

import (
	"at3-back/config"
	"at3-back/internal/shared/infrastructure/entrypoint/http_server/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	app.Use(recover.New())
	app.Use(healthcheck.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://app.atomico3.io, http://localhost:8080, http://localhost:8081",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowHeaders: "Origin, Accept, Authorization, Content-Type, X-CSRF-Token",
	}))

	routes.Init(app)

	return app.Listen(os.Getenv("PORT"))
}
