package routes

import (
	"test-dbo/controllers"
	"test-dbo/middleware"
	"test-dbo/repositories"
	"test-dbo/services"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	// Initialize repository & service
	orderRepo := repositories.NewOrderRepository()
	orddrService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orddrService)

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/order", orderController.CreateOrder)
		api.GET("/orders", orderController.GetAllOrders)
		api.GET("/order/:id", orderController.GetOrderByID)
	}
}
