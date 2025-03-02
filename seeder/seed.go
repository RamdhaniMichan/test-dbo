package seeder

import (
	"fmt"

	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	fmt.Println("Seeding database...")

	users := SeedUsers(db)
	products := SeedProducts(db)
	orders := SeedOrders(db, users)
	SeedOrderItems(db, orders, products)

	fmt.Println("Seeding completed!")
}
