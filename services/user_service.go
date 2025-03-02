package services

import (
	"errors"
	"test-dbo/models"
	"test-dbo/repositories"
	"test-dbo/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers(pagination utils.Pagination) ([]models.User, int64, error)
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	DeleteUserByID(id uint) error
	UpdateUserByID(id uint, updates map[string]interface{}) error
	GetUserHistory(id uint) ([]models.UserHistory, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{userRepo: repo}
}

// GetUserHistory implements UserService.
func (s *userService) GetUserHistory(id uint) ([]models.UserHistory, error) {
	return s.userRepo.GetUserHistory(id)
}

// UpdateUserByID implements services.UserService.
func (s *userService) UpdateUserByID(id uint, updates map[string]interface{}) error {
	return s.userRepo.Update(id, updates)
}

// DeleteUserByID implements services.UserService.
func (s *userService) DeleteUserByID(id uint) error {
	return s.userRepo.Delete(id)
}

// GetUserByID implements services.UserService.
func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

// Get all users
func (s *userService) GetUsers(pagination utils.Pagination) ([]models.User, int64, error) {
	return s.userRepo.GetAll(pagination)
}

// Create new user with validation
func (s *userService) CreateUser(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.Create(user)
}
