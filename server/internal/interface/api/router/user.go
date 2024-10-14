package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/real-estate-admin/server/internal/interface/api/handler"
)

func SetupUserRouter(app fiber.Router, h *handler.User) {
	app.Get("/users", h.List)
	app.Get("/users/:id", h.Get)
	app.Post("/users", h.Create)
}
