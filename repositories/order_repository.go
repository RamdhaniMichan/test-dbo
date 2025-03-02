package repositories

import (
	"test-dbo/database"
	"test-dbo/models"
	"test-dbo/utils"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	CreateOrderItem(item *models.OrderItem) error
	UpdateProductStock(productID uint, quantity int) error
	GetAllOrders(pagination utils.Pagination, role string, user_id uint) ([]models.Order, int64, error)
	GetOrderByID(orderID uint) (models.Order, error)
}

type orderRepository struct{}

// GetOrderByID implements OrderRepository.
func (o *orderRepository) GetOrderByID(orderID uint) (models.Order, error) {
	var order models.Order
	err := database.DB.Preload("User").Preload("Items.Product").First(&order, orderID).Error
	return order, err
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

// GetAllOrders implements OrderRepository.
func (o *orderRepository) GetAllOrders(pagination utils.Pagination, role string, user_id uint) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	query := database.DB.Model(&models.Order{})

	if status, exists := pagination.Filters["status"]; exists {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if role != "admin" {
		query = query.Where("user_id = ?", user_id)
		query.Count(&total)
	}

	err := query.Preload("User").Preload("Items.Product").
		Limit(pagination.Limit).Offset(pagination.Offset).Find(&orders).Error
	return orders, total, err
}

// CreateOrder implements OrderRepository.
func (o *orderRepository) CreateOrder(order *models.Order) error {
	return database.DB.Create(order).Error
}

// CreateOrderItem implements OrderRepository.
func (o *orderRepository) CreateOrderItem(item *models.OrderItem) error {
	return database.DB.Create(item).Error
}

// UpdateProductStock implements OrderRepository.
func (o *orderRepository) UpdateProductStock(productID uint, quantity int) error {
	return database.DB.Model(&models.Product{}).
		Where("id = ?", productID).
		Update("stock", gorm.Expr("stock - ?", quantity)).Error
}
