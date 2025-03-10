package repositories

import (
	"github.com/bps-kota-bontang/fiber-starter/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(id uint) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) error
}
