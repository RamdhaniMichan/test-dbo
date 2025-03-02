package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware RBAC (Role-Based Access Control)
func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}
