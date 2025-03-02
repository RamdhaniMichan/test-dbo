package repositories

import (
	"test-dbo/database"
	"test-dbo/models"
	"test-dbo/utils"
)

type UserRepository interface {
	GetAll(pagination utils.Pagination) ([]models.User, int64, error)
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	Delete(id uint) error
	Update(id uint, updates map[string]interface{}) error
	GetUserHistory(id uint) ([]models.UserHistory, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// GetUserHistory implements UserRepository.
func (r *userRepository) GetUserHistory(id uint) ([]models.UserHistory, error) {
	panic("unimplemented")
}

// UpdateUserByID implements UserRepository.
func (r *userRepository) Update(id uint, updates map[string]interface{}) error {
	return database.DB.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteByID implements UserRepository.
func (r *userRepository) Delete(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}

// GetByID implements UserRepository.
func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

// Get all users
func (r *userRepository) GetAll(pagination utils.Pagination) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := database.DB.Model(&models.User{})

	if email, exists := pagination.Filters["email"]; exists {
		query = query.Where("email ILIKE ?", "%"+email+"%")
	}

	query.Count(&total)

	err := query.Limit(pagination.Limit).Offset(pagination.Offset).Find(&users).Error
	return users, total, err
}

// Create new user
func (r *userRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}
