package app

import (
	"github.com/bps-kota-bontang/fiber-starter/config"
	"github.com/bps-kota-bontang/fiber-starter/handlers"
	"github.com/bps-kota-bontang/fiber-starter/routes"

	"github.com/gofiber/fiber/v2"
)

// NewFiberApp initializes the Fiber application with all necessary routes
func NewFiberApp(
	appConfig config.AppConfig,
	userHandler handlers.UserHandler,
) *fiber.App {
	app := fiber.New(
		fiber.Config{
			AppName: appConfig.AppName,
		},
	)

	// API Versioning
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Register routes
	routes.RegisterUserRoutes(v1, userHandler)

	return app
}
