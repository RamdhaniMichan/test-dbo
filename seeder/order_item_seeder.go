package seeder

import (
	"test-dbo/models"

	"gorm.io/gorm"
)

func SeedOrderItems(db *gorm.DB, orders []models.Order, products []models.Product) {
	orderItems := []models.OrderItem{
		{OrderID: orders[0].ID, ProductID: products[0].ID, Quantity: 1, Price: 1500.00},
		{OrderID: orders[0].ID, ProductID: products[1].ID, Quantity: 1, Price: 800.00},
		{OrderID: orders[1].ID, ProductID: products[0].ID, Quantity: 1, Price: 1500.00},
	}

	db.Create(&orderItems)
}
