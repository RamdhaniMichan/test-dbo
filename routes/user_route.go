package routes

import (
	"test-dbo/controllers"
	"test-dbo/middleware"
	"test-dbo/repositories"
	"test-dbo/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	// Initialize repository & service
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
	{
		api.GET("/users", userController.GetUsers)
		api.POST("/users", userController.CreateUser)
		api.GET("/user/:id", userController.GetByID)
		api.DELETE("/user/:id", userController.DeleteCustomer)
		api.PUT("/user/:id", userController.UpdateCustomer)
	}
}
