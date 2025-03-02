package seeder

import (
	"test-dbo/models"

	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) []models.Product {
	products := []models.Product{
		{Name: "Laptop", Description: "Gaming laptop", Price: 1500.00, Stock: 10},
		{Name: "Smartphone", Description: "Android phone", Price: 800.00, Stock: 20},
	}

	db.Create(&products)
	return products
}
