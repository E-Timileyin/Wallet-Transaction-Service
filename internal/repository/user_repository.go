package repository

import (
	"wallet-service/internal/db"
	"wallet-service/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{DB: db.DB}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}

func (r *UserRepository) DeleteAll() error {
	return r.DB.Exec("DELETE FROM users").Error
}
