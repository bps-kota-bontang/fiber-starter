package routes

import (
	"github.com/bps-kota-bontang/fiber-starter/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router, handler handlers.UserHandler) {
	user := router.Group("/users")
	user.Get("/", handler.GetUsers)
	user.Post("/", handler.CreateUser)
	user.Get("/:id", handler.GetUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)
}
