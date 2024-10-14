package server

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gotired/real-estate-admin/server/internal/config"
	"github.com/gotired/real-estate-admin/server/internal/domain/model"
	"github.com/gotired/real-estate-admin/server/internal/infrastructure/middleware"
	"github.com/gotired/real-estate-admin/server/internal/infrastructure/persistance/database"
	_ "github.com/lib/pq"
)

func Setup() (*fiber.App, *model.CONFIG, *sql.DB) {

	Config := config.Init()
	db := database.Init(&Config.DB)

	app := fiber.New()
	app.Use(middleware.LogRequest)

	return app, Config, db
}
