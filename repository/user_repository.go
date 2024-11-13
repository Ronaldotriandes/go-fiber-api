package repository

import (
	"github.com/Ronaldotriandes/go-fiber-api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetById(id uint) (*models.User, error)
	GetAll() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}
func (r *userRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}
