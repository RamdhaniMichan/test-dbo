package routes

import (
	"test-dbo/controllers"
	"test-dbo/middleware"
	"test-dbo/repositories"
	"test-dbo/services"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	// Initialize repository & service
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
	{
		api.POST("/product", productController.CreateProduct)
		api.GET("/products", productController.GetProducts)
		api.GET("/product/:id", productController.GetProductByID)
		api.PUT("/product/:id", productController.UpdateProduct)
		api.DELETE("/product/:id", productController.DeleteProduct)
	}
}
