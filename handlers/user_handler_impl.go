package handlers

import (
	"github.com/bps-kota-bontang/fiber-starter/dto"
	"github.com/bps-kota-bontang/fiber-starter/errors"
	"github.com/bps-kota-bontang/fiber-starter/services"
	"github.com/bps-kota-bontang/fiber-starter/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerImpl struct {
	service  services.UserService
	validate *validator.Validate
}

func NewUserHandler(service services.UserService, validate *validator.Validate) UserHandler {
	return &UserHandlerImpl{
		service:  service,
		validate: validate,
	}
}

// CreateUser implements UserHandler.
func (b *UserHandlerImpl) CreateUser(c *fiber.Ctx) error {

	var user dto.CreateUserRequest
	if err := c.BodyParser(&user); err != nil {
		errorResponse, code := utils.HandleErrors(errors.NewHttpError(fiber.StatusBadRequest, "Invalid request body"))
		return c.Status(code).JSON(errorResponse)
	}

	if err := b.validate.Struct(user); err != nil {
		errorResponse, code := utils.HandleErrors(err)
		return c.Status(code).JSON(errorResponse)
	}

	userCreated, err := b.service.CreateUser(&user)
	if err != nil {
		errorResponse, code := utils.HandleErrors(err)
		return c.Status(code).JSON(errorResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    userCreated,
		"message": "User created successfully",
	})
}

// DeleteUser implements UserHandler.
func (b *UserHandlerImpl) DeleteUser(c *fiber.Ctx) error {
	id, err := utils.ParseID(c.Params("id"))
	if err != nil {
		errorResponse, code := utils.HandleErrors(errors.NewHttpError(fiber.StatusBadRequest, "Invalid user ID"))
		return c.Status(code).JSON(errorResponse)
	}

	if err := b.service.DeleteUser(id); err != nil {
		errorResponse, code := utils.HandleErrors(err)
		return c.Status(code).JSON(errorResponse)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetUser implements UserHandler.
func (b *UserHandlerImpl) GetUser(c *fiber.Ctx) error {
	id, err := utils.ParseID(c.Params("id"))
	if err != nil {
		errorResponse, code := utils.HandleErrors(errors.NewHttpError(fiber.StatusBadRequest, "Invalid user ID"))
		return c.Status(code).JSON(errorResponse)
	}

	user, err := b.service.GetUserById(id)
	if err != nil {
		errorResponse, code := utils.HandleErrors(err)
		return c.Status(code).JSON(errorResponse)
	}

	return c.JSON(fiber.Map{
		"data":    user,
		"message": "User retrieved successfully",
	})
}

// GetUsers implements UserHandler.
func (b *UserHandlerImpl) GetUsers(c *fiber.Ctx) error {
	users, err := b.service.GetAllUsers()

	if err != nil {
		errorResponse, code := utils.HandleErrors(err)
		return c.Status(code).JSON(errorResponse)
	}

	return c.JSON(fiber.Map{
		"data":    users,
		"message": "Users retrieved successfully",
	})
}

// UpdateUser implements UserHandler.
func (b *UserHandlerImpl) UpdateUser(c *fiber.Ctx) error {
	var user dto.UpdateUserRequest

	id, err := utils.ParseID(c.Params("id"))
	if err != nil {
		errorResponse, code := utils.HandleErrors(errors.NewHttpError(fiber.StatusBadRequest, "Invalid user ID"))
		return c.Status(code).JSON(errorResponse)
	}

	if err := c.BodyParser(&user); err != nil {
		errorResponse, code := utils.HandleErrors(errors.NewHttpError(fiber.StatusBadRequest, "Invalid request body"))
		return c.Status(code).JSON(errorResponse)
	}

	if err := b.validate.Struct(user); err != nil {
		errorResponse, code := utils.HandleErrors(err)
		return c.Status(code).JSON(errorResponse)
	}

	userUpdated, err := b.service.UpdateUser(id, &user)

	if err != nil {
		errorResponse, code := utils.HandleErrors(err)
		return c.Status(code).JSON(errorResponse)
	}

	return c.JSON(fiber.Map{
		"data":    userUpdated,
		"message": "User updated successfully",
	})
}
