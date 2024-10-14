package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/real-estate-admin/server/internal/interface/api/handler"
)

func SetupPropertyRouter(app fiber.Router, h *handler.Property) {
	app.Get("/properties", h.List)
	app.Get("/properties/:id", h.Get)
	app.Post("/properties", h.Create)
	app.Patch("/properties/:id", h.Update)
	app.Delete("/properties/:id", h.Delete)
}
