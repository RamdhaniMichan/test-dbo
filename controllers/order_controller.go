package controllers

import (
	"net/http"
	"strconv"
	"test-dbo/models"
	"test-dbo/services"
	"test-dbo/utils"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{orderService: service}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var request struct {
		UserID uint               `json:"user_id"`
		Items  []models.OrderItem `json:"items"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := c.orderService.CreateOrder(request.UserID, request.Items)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// Get all orders
func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	role, _ := ctx.Get("role")

	// Convert userID ke uint
	uid, ok := userID.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	pagination := utils.Paginate(ctx, []string{"status"})
	orders, total, err := c.orderService.GetAllOrders(pagination, role.(string), uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       orders,
		"total_data": total,
		"page":       pagination.Page,
		"limit":      pagination.Limit,
	})
}

// Get order by ID
func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := c.orderService.GetOrderByID(uint(orderID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}
