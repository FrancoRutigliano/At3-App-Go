package routes

import (
	authRoutes "at3-back/internal/auth/infrastructure/routes"

	"github.com/gofiber/fiber/v2"
)

func Init(f *fiber.App) {
	api := f.Group("/api/v1")

	authRoutes.Init(api)
}
