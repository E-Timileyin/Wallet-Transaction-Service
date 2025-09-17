package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// temporary in-memory users
var users = []User{
	{Id: 1, Name: "John Doe", Email: "john.doe@example.com", Password: "password123"},
	{Id: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Password: "password456"},
}

// GET /users (all users or query by email)
func GetUsers(c *gin.Context) {
	email := c.Query("email")
	if email != "" {
		for _, user := range users {
			if user.Email == email {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GET /users/:id (user by ID)
func GetUserById(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if strconv.Itoa(user.Id) == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// POST /users (create new user)
func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// naive auto-increment ID
	user.Id = len(users) + 1
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}
