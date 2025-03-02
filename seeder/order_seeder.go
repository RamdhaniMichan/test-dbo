package seeder

import (
	"test-dbo/models"

	"gorm.io/gorm"
)

func SeedOrders(db *gorm.DB, users []models.User) []models.Order {
	orders := []models.Order{
		{UserID: users[0].ID, Total: 2300.00, Status: "pending"},
		{UserID: users[1].ID, Total: 1500.00, Status: "shipped"},
	}

	db.Create(&orders)
	return orders
}
