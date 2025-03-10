package services

import (
	"github.com/bps-kota-bontang/fiber-starter/dto"
	"github.com/bps-kota-bontang/fiber-starter/models"
	"github.com/bps-kota-bontang/fiber-starter/repositories"
	"github.com/bps-kota-bontang/fiber-starter/utils"
)

type UserServiceImpl struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &UserServiceImpl{repository}
}

func (b *UserServiceImpl) CreateUser(user *dto.CreateUserRequest) (*dto.UserResponse, error) {
	newUser := &models.User{
		Name: user.Name,
	}

	userCreated, err := b.repository.Create(newUser)

	if err != nil {
		return nil, err
	}

	response := utils.ToUserResponse(userCreated)
	return &response, nil
}

func (b *UserServiceImpl) DeleteUser(id uint) error {
	_, err := b.repository.FindById(id)
	if err != nil {
		return err
	}

	return b.repository.Delete(id)
}

func (b *UserServiceImpl) GetAllUsers() ([]dto.UserResponse, error) {
	users, err := b.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return utils.ToUserResponses(users), nil
}

func (b *UserServiceImpl) GetUserById(id uint) (*dto.UserResponse, error) {
	user, err := b.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	response := utils.ToUserResponse(user)
	return &response, nil
}

func (b *UserServiceImpl) UpdateUser(id uint, user *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	userToUpdate, err := b.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	if user.Name != nil {
		userToUpdate.Name = *user.Name
	}

	userUpdated, err := b.repository.Update(userToUpdate)
	if err != nil {
		return nil, err
	}

	response := utils.ToUserResponse(userUpdated)
	return &response, nil
}
