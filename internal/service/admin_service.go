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

// UpdateUserRole updates the role of a specific user.
func (s *AdminService) UpdateUserRole(userID uint, newRole string) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	user.Role = newRole
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteAllUsers deletes all users from the system.
func (s *AdminService) DeleteAllUsers() error {
	return s.userRepo.DeleteAll()
}
