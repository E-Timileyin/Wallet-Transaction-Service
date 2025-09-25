package handlers

import (
	"net/http"
	"strconv"
	"wallet-service/internal/models"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
)

// AdminHandler handles admin-related HTTP requests.
type AdminHandler struct {
	adminService *service.AdminService
}

// NewAdminHandler creates a new AdminHandler.
func NewAdminHandler(adminService *service.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

// GetAllUsers handles the request to get all users.
func (h *AdminHandler) GetAllUsers(c *gin.Context) {
	users, err := h.adminService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUserRole handles the request to update a user's role.
func (h *AdminHandler) UpdateUserRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req models.UpdateUserRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.adminService.UpdateUserRole(uint(id), req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteAllUsers handles the request to delete all users.
func (h *AdminHandler) DeleteAllUsers(c *gin.Context) {
	if err := h.adminService.DeleteAllUsers(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All users deleted successfully"})
}
