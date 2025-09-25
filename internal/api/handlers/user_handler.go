package handlers

import (
	"net/http"
	"strconv"
	"wallet-service/internal/models"
	"wallet-service/internal/service"

	"github.com/gin-gonic/gin"
	_ "wallet-service/docs"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}

// GetUsers retrieves all users
// @Summary Get all users
// @Description Retrieve a list of all users in the system
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /api/users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a user by ID
// @Summary Get user by ID
// @Description Retrieve a specific user by their ID
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.Service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUserByEmail retrieves a user by email
// @Summary Get user by email
// @Description Retrieve a specific user by their email address
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /api/users/email/{email} [get]
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.Service.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user account
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

