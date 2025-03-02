package services

import (
	"test-dbo/models"
	"test-dbo/repositories"
	"test-dbo/utils"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProducts(pagination utils.Pagination) ([]models.Product, int64, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(id uint, updates map[string]interface{}) error
	DeleteProduct(id uint) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repoProduct repositories.ProductRepository) ProductService {
	return &productService{repo: repoProduct}
}

// CreateProduct implements ProductService.
func (p *productService) CreateProduct(product *models.Product) error {
	return p.repo.Create(product)
}

// DeleteProduct implements ProductService.
func (p *productService) DeleteProduct(id uint) error {
	return p.repo.Delete(id)
}

// GetProductByID implements ProductService.
func (p *productService) GetProductByID(id uint) (*models.Product, error) {
	return p.repo.FindByID(id)
}

// GetProducts implements ProductService.
func (p *productService) GetProducts(pagination utils.Pagination) ([]models.Product, int64, error) {
	return p.repo.FindAll(pagination)
}

// UpdateProduct implements ProductService.
func (p *productService) UpdateProduct(id uint, updates map[string]interface{}) error {
	return p.repo.Update(id, updates)
}
