package di

import (
	"github.com/bps-kota-bontang/fiber-starter/config"
	"github.com/gofiber/fiber/v2"
)

type AppContainer struct {
	App    *fiber.App
	Config config.AppConfig
}
