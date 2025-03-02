package models

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `json:"order_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"`

	Order   Order   `gorm:"foreignKey:OrderID"`
	Product Product `gorm:"foreignKey:ProductID"`
}
