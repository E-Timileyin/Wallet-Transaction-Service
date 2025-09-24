package service

import (
	"errors"
	"time"
	"wallet-service/internal/config"
	"wallet-service/internal/models"
	"wallet-service/internal/repository"
	"wallet-service/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// AuthService provides authentication-related services.
type AuthService struct {
	userRepo repository.UserRepository
	cfg      *config.Config
}

// NewAuthService creates a new AuthService.
func NewAuthService(userRepo repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{userRepo: userRepo, cfg: cfg}
}

// Register creates a new user, hashes their password, and saves them to the database.
func (s *AuthService) Register(req *models.RegisterRequest) (*models.User, error) {
	// Check if user already exists
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // The password will be hashed by the BeforeCreate hook
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Do not return the password hash
	user.Password = ""
	return user, nil
}

// Login authenticates a user and returns an authentication response with JWT tokens.
func (s *AuthService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := s.generateToken(user, s.cfg.JWT.Expiration)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateToken(user, s.cfg.JWT.RefreshExpiration)
	if err != nil {
		return nil, err
	}

	// Do not return the password hash
	user.Password = ""

	return &models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

// RefreshToken generates a new access token from a valid refresh token.
func (s *AuthService) RefreshToken(refreshTokenString string) (*models.AuthResponse, error) {
	claims, err := s.ValidateToken(refreshTokenString)
	if err != nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	userID := uint(claims["user_id"].(float64))
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	accessToken, err := s.generateToken(user, s.cfg.JWT.Expiration)
	if err != nil {
		return nil, err
	}

	// Do not return the password hash
	user.Password = ""

	return &models.AuthResponse{
		AccessToken: accessToken,
		User:        user,
	}, nil
}

// ValidateToken parses and validates a JWT token string.
func (s *AuthService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// generateToken creates a new JWT token for a user.
func (s *AuthService) generateToken(user *models.User, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(expiration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWT.Secret))
}
