package models

import "time"

type UserHistory struct {
	ID         uint                   `gorm:"primaryKey"`
	CustomerID uint                   `gorm:"index"`
	Changes    map[string]interface{} `gorm:"type:jsonb"`
	UpdatedAt  time.Time
}
