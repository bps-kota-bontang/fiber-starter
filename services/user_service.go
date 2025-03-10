package services

import (
	"github.com/bps-kota-bontang/fiber-starter/dto"
)

type UserService interface {
	GetAllUsers() ([]dto.UserResponse, error)
	GetUserById(id uint) (*dto.UserResponse, error)
	CreateUser(user *dto.CreateUserRequest) (*dto.UserResponse, error)
	UpdateUser(id uint, user *dto.UpdateUserRequest) (*dto.UserResponse, error)
	DeleteUser(id uint) error
}
