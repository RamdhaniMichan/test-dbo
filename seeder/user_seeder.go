package seeder

import (
	"test-dbo/models"

	"gorm.io/gorm"
)

// Ganti dengan path models yang sesuai

func SeedUsers(db *gorm.DB) []models.User {
	// default_password = 12345678
	users := []models.User{
		{Name: "John Doe", Email: "john@example.com", Password: "$2a$10$o56o21tTvMO23I.csoQPKuLRwQ51pJNe9.YYjqmDU10M07uA90rGG", PhoneNumber: "123456789", Address: "New York", BirthDate: "1990-01-01", Gender: "Male", Status: "active", Role: "customer"},
		{Name: "Jane Smith", Email: "jane@example.com", Password: "$2a$10$o56o21tTvMO23I.csoQPKuLRwQ51pJNe9.YYjqmDU10M07uA90rGG", PhoneNumber: "987654321", Address: "California", BirthDate: "1992-05-12", Gender: "Female", Status: "active", Role: "customer"},
	}

	db.Create(&users)
	return users
}
