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

	response := a.handler.Impl.Confirm(token)
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "data": response.Data, "details": "false"})
	}

	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "data": response.Data, "details": "true"})
}
