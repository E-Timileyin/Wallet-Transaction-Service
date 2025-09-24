package service

import (
	"wallet-service/internal/models"
	"wallet-service/internal/repository"
)

// AdminService provides admin-related services.
type AdminService struct {
	userRepo *repository.UserRepository
}

// NewAdminService creates a new AdminService.
func NewAdminService(userRepo *repository.UserRepository) *AdminService {
	return &AdminService{userRepo: userRepo}
}

// GetAllUsers retrieves all users from the repository.
func (s *AdminService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}
