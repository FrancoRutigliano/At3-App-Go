package authController

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	"at3-back/pkg/validator"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Auth) Register(c *fiber.Ctx) error {
	var payload authDto.RegisterUser

	response := validator.Payload(c, &payload)
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}

	response = a.handler.Impl.Register(payload)
	if response.StatusCode != http.StatusCreated {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}

	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "data": response.Data, "details": "true"})
}
