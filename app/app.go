package app

import (
	"github.com/bps-kota-bontang/fiber-starter/handlers"
	"github.com/bps-kota-bontang/fiber-starter/routes"

	"github.com/gofiber/fiber/v2"
)

// NewFiberApp initializes the Fiber application with all necessary routes
func NewFiberApp(
	userHandler handlers.UserHandler,
) *fiber.App {
	app := fiber.New()

	// API Versioning
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Register routes
	routes.RegisterUserRoutes(v1, userHandler)

	return app
}
