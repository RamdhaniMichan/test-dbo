package services

import (
	"errors"
	"test-dbo/database"
	"test-dbo/models"
	"test-dbo/repositories"
	"test-dbo/utils"
)

type OrderService interface {
	CreateOrder(userID uint, items []models.OrderItem) (models.Order, error)
	GetAllOrders(pagination utils.Pagination, role string, user_id uint) ([]models.Order, int64, error)
	GetOrderByID(orderID uint) (models.Order, error)
}

type orderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(repoOrder repositories.OrderRepository) OrderService {
	return &orderService{repo: repoOrder}
}

// GetOrderByID implements OrderService.
func (o *orderService) GetOrderByID(orderID uint) (models.Order, error) {
	return o.repo.GetOrderByID(orderID)
}

// GetAllOrders implements OrderService.
func (o *orderService) GetAllOrders(pagination utils.Pagination, role string, user_id uint) ([]models.Order, int64, error) {
	return o.repo.GetAllOrders(pagination, role, user_id)
}

// CreateOrder implements OrderService.
func (o *orderService) CreateOrder(userID uint, items []models.OrderItem) (models.Order, error) {
	var total float64

	if userID == 0 {
		return models.Order{}, errors.New("user_id is required")
	}

	if len(items) == 0 {
		return models.Order{}, errors.New("order must have at least one item")
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return models.Order{}, errors.New("user not found")
	}

	order := models.Order{
		UserID: userID,
		Status: "Pending",
	}

	tx := database.DB.Begin()
	if err := o.repo.CreateOrder(&order); err != nil {
		tx.Rollback()
		return models.Order{}, err
	}

	for i := range items {
		var product models.Product
		if err := tx.First(&product, items[i].ProductID).Error; err != nil {
			tx.Rollback()
			return models.Order{}, errors.New("product not found")
		}

		if items[i].Quantity > product.Stock {
			tx.Rollback()
			return models.Order{}, errors.New("insufficient stock for product " + product.Name)
		}
	}

	for i := range items {
		items[i].OrderID = order.ID
		total += items[i].Price * float64(items[i].Quantity)
		if err := o.repo.CreateOrderItem(&items[i]); err != nil {
			tx.Rollback()
			return models.Order{}, err
		}
		if err := o.repo.UpdateProductStock(items[i].ProductID, items[i].Quantity); err != nil {
			tx.Rollback()
			return models.Order{}, err
		}
	}

	order.Total = total
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return models.Order{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return models.Order{}, err
	}
	return order, nil
}
