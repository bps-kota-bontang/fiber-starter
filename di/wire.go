//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// Initialize App
func InitializeApp() (*fiber.App, error) {
	wire.Build(AppSet)
	return &fiber.App{}, nil
}
