package authController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Auth) Confirm_account(c *fiber.Ctx) error {

	token := c.Query("token")
	if token == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "token required", "details": "false"})
	}

	entity := c.Query("type")
	if entity != "user" && entity != "company" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid token", "details": "false"})
	}

	response := a.handler.Impl.Confirm(token, entity)
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}

	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "data": response.Data, "details": "true"})
}
