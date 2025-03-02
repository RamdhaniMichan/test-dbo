package main

import (
	"fmt"
	"test-dbo/config"
	"test-dbo/database"
	"test-dbo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config.LoadEnv()

	// Initialize database
	database.ConnectDB()

	// Setup Routes
	r := gin.Default()
	routes.UserRoutes(r)
	routes.ProductRoutes(r)
	routes.OrderRoutes(r)
	routes.AuthRoutes(r)

	// Run server
	port := config.GetEnv("PORT", "8080")
	r.Run(fmt.Sprintf(":%s", port))
}
