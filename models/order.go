package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID uint    `json:"user_id" gorm:"not null"`
	Total  float64 `json:"total" gorm:"not null"`
	Status string  `json:"status" gorm:"type:varchar(20);not null"`

	User  User        `gorm:"foreignKey:UserID"`
	Items []OrderItem `gorm:"foreignKey:OrderID"`
}
