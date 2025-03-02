package repositories

import (
	"errors"
	"test-dbo/database"
	"test-dbo/models"
)

type AuthRepository interface {
	Register(user *models.User) error
	Login(email string) (*models.User, error)
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

// Register implements AuthRepository.
func (a *authRepository) Register(user *models.User) error {
	return database.DB.Create(user).Error
}

// Login implements AuthRepository.
func (a *authRepository) Login(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
