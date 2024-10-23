package validator

import (
	httpresponse "at3-back/pkg/httpResponse"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Payload(c *fiber.Ctx, payload interface{}) httpresponse.ApiResponse {
	if err := c.BodyParser(payload); err != nil {
		return *httpresponse.NewApiError(http.StatusBadRequest, "invalid request", nil)
	}

	return httpresponse.ApiResponse{}
}
