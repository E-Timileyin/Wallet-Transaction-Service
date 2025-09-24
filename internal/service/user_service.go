package service

import (
	"wallet-service/internal/models"
	"wallet-service/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.FindByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.Repo.FindByEmail(email)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.Repo.Create(user)
}
func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.Delete(id)
}
