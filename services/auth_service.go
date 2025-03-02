package services

import (
	"errors"
	"test-dbo/models"
	"test-dbo/repositories"
	"test-dbo/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user *models.User) (string, error)
	Login(email, password string) (string, error)
}

type authService struct {
	repo repositories.AuthRepository
}

// Register implements AuthService.
func (a *authService) Register(user *models.User) (string, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	err = a.repo.Register(user)
	if err != nil {
		return "", err
	}

	// Generate JWT
	return utils.GenerateToken(user.ID, user.Role)
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authService{repo: repository}
}

// Login implements AuthService.
func (a *authService) Login(email string, password string) (string, error) {
	//Get user
	user, err := a.repo.Login(email)
	if err != nil {
		return "", err
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT
	return utils.GenerateToken(user.ID, user.Role)
}
