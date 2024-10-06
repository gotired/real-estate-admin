package middleware

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gotired/real-estate-admin/server/internal/interface/api/response"
)

func LogRequest(c *fiber.Ctx) error {
	start := time.Now()

	if err := c.Next(); err != nil {
		if strings.HasPrefix(err.Error(), "Cannot") {
			logRequest(c, start, fiber.StatusNotFound)
			return response.Error(c, fiber.StatusNotFound, err.Error())
		}
		c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		logRequest(c, start, fiber.StatusInternalServerError)
		return err
	}

	logRequest(c, start, c.Response().StatusCode())
	return nil
}

func logRequest(c *fiber.Ctx, start time.Time, statusCode int) {
	resSize := len(c.Response().Body())
	responseTime := time.Since(start)

	log.Printf("[%s] - %s - %s %s %d %d %.2f ms",
		time.Now().Format("02/Jan/2006:15:04:05 -0700"), c.IP(), c.Method(), c.Path(),
		statusCode, resSize, responseTime.Seconds()*1000)
}
