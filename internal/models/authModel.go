package models

// RegisterRequest defines the structure for a user registration request.
// @Description User registration request
type RegisterRequest struct {
	Name     string `json:"name" example:"John Doe" binding:"required"`
	Email    string `json:"email" example:"john@example.com" binding:"required,email"`
	Password string `json:"password" example:"password123" binding:"required,min=8"`
}

// LoginRequest defines the structure for a user login request.
// @Description User login request
type LoginRequest struct {
	Email    string `json:"email" example:"john@example.com" binding:"required,email"`
	Password string `json:"password" example:"password123" binding:"required"`
}

// AuthResponse defines the structure for a successful authentication response.
// @Description Authentication response
type AuthResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	User         *User  `json:"user"`
}

// RefreshTokenRequest defines the structure for a token refresh request.
// @Description Token refresh request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" binding:"required"`
}
