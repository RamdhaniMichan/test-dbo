package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-dbo/database"
	"test-dbo/models"
	"test-dbo/services"
	"test-dbo/utils"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{
		productService: service,
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pc.productService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	pagination := utils.Paginate(c, []string{"name"})
	products, total, err := pc.productService.GetProducts(pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":       products,
		"total_data": total,
		"page":       pagination.Page,
		"limit":      pagination.Limit,
	})
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	product, err := pc.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := database.DB.First(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(c.Request.Body).Decode(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := pc.productService.UpdateProduct(uint(id), updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, updates)
}

// Delete Product By ID
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = pc.productService.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
