package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/real-estate-admin/server/internal/domain/model"
)

func Success(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(model.Response{
		Status: "success",
		Detail: "",
		Data:   data,
	})
}

func Error(c *fiber.Ctx, statusCode int, detail string) error {
	return c.Status(statusCode).JSON(model.Response{
		Status: "failure",
		Detail: detail,
		Data:   nil,
	})
}
