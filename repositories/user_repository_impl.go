package repositories

import (
	"github.com/bps-kota-bontang/fiber-starter/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	var user models.User

	return r.db.Delete(&user, id).Error
}

func (r *UserRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindById(id uint) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
