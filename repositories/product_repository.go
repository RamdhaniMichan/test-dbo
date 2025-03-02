package repositories

import (
	"test-dbo/database"
	"test-dbo/models"
	"test-dbo/utils"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindAll(pagination utils.Pagination) ([]models.Product, int64, error)
	FindByID(id uint) (*models.Product, error)
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

// Create implements ProductRepository.
func (p *productRepository) Create(product *models.Product) error {
	return database.DB.Create(product).Error
}

// Delete implements ProductRepository.
func (p *productRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Product{}, id).Error
}

// FindAll implements ProductRepository.
func (p *productRepository) FindAll(pagination utils.Pagination) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := database.DB.Model(&models.Product{})

	if name, exists := pagination.Filters["name"]; exists {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	query.Count(&total)

	err := query.Limit(pagination.Limit).Offset(pagination.Offset).Find(&products).Error
	return products, total, err
}

// FindByID implements ProductRepository.
func (p *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := database.DB.First(&product, id).Error
	return &product, err
}

// Update implements ProductRepository.
func (p *productRepository) Update(id uint, updates map[string]interface{}) error {
	return database.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updates).Error
}
