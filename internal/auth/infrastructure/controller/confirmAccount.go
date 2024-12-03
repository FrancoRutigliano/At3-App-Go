package authController

import (
	httpresponse "at3-back/pkg/httpResponse"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Auth) Confirm_account(c *fiber.Ctx) error {
	var response httpresponse.ApiResponse
	token := c.Query("token")
	if token == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "token required", "details": "false"})
	}

	entity := c.Query("type")

	if entity == "user" {
		response = a.handler.Impl.ConfirmUserAccount(token, entity)
		if response.StatusCode != http.StatusOK {
			return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
		}

	} else if entity == "company" {

	} else {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid entity", "details": "false"})
	}

	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "true", "data": response.Data})
}
