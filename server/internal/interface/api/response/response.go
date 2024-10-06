package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/real-estate-admin/server/internal/domain/model"
)

func Success(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: "success",
		Detail: "",
		Data:   data,
	})
}

func Created(c *fiber.Ctx, data interface{}, detail string) error {
	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Status: "success",
		Detail: detail,
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
