package routes

import (
	"test-dbo/controllers"
	"test-dbo/repositories"
	"test-dbo/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	// Initialize repository & service
	authRepo := repositories.NewAuthRepository()
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	api := router.Group("/api")
	{
		api.POST("/login", authController.Login)
		api.POST("/register", authController.Register)
	}
}
