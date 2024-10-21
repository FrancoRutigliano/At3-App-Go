package authRoutes

import "github.com/gofiber/fiber/v2"

func Init(r fiber.Router) {

	r.Post("/login")
	r.Post("/register")
	r.Post("/reset")

}
