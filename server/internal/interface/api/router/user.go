package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/real-estate-admin/server/internal/interface/api/handler"
)

func SetupUserRouter(app fiber.Router, userHandler *handler.User) {
	app.Get("/users", userHandler.List)
	app.Get("/users/:id", userHandler.Get)
	app.Post("/users", userHandler.Create)
}
