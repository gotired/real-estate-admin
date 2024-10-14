package main

import (
	"fmt"
	"log"

	"github.com/gotired/real-estate-admin/server/internal/domain/service"
	"github.com/gotired/real-estate-admin/server/internal/infrastructure/persistance/property"
	"github.com/gotired/real-estate-admin/server/internal/infrastructure/persistance/user"
	"github.com/gotired/real-estate-admin/server/internal/infrastructure/server"
	"github.com/gotired/real-estate-admin/server/internal/interface/api/handler"
	"github.com/gotired/real-estate-admin/server/internal/interface/api/router"
)

func main() {
	app, config, db := server.Setup()

	userRepo := user.New(db)
	userService := &service.User{Repository: userRepo}

	propertyRepo := property.New(db)
	propertyService := &service.Property{Repository: propertyRepo}

	userHandler := &handler.User{UserService: userService, PropertyService: propertyService}
	propertyHandler := &handler.Property{UserService: userService, PropertyService: propertyService}

	apiRouter := app.Group("/api")
	v1Router := apiRouter.Group("/v1")

	router.SetupUserRouter(v1Router, userHandler)
	router.SetupPropertyRouter(v1Router, propertyHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.APP.Port)))
}
