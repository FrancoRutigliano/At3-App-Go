package validator

import (
	httpresponse "at3-back/pkg/httpResponse"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Payload(c *fiber.Ctx, payload interface{}) httpresponse.ApiResponse {
	if err := c.BodyParser(payload); err != nil {
		log.Println("Failed to parse body request: Error:", err)
		return *httpresponse.NewApiError(http.StatusBadRequest, "invalid request", nil)
	}

	return httpresponse.ApiResponse{}
}

func DateToUnix() int64 {
	now := time.Now()

	_ = now.Format("02:01:2006")

	return now.Unix()
}
