package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(100);not null"`
	Description string  `json:"description" gorm:"type:text"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock" gorm:"not null"`
}

// Validate function for Product
func (p Product) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&p.Price, validation.Required, validation.Min(0.01)), // Minimal harga 0.01
		validation.Field(&p.Stock, validation.Required, validation.Min(0)),    // Stok minimal 0
	)
}
