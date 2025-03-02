package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination struct
type Pagination struct {
	Limit   int               `json:"limit"`
	Page    int               `json:"page"`
	Offset  int               `json:"offset"`
	Filters map[string]string `json:"filters"`
}

// Paginate function
func Paginate(ctx *gin.Context, filterKeys []string) Pagination {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	filters := make(map[string]string)
	for _, key := range filterKeys {
		if value := ctx.Query(key); value != "" {
			filters[key] = value
		}
	}

	return Pagination{
		Limit:   limit,
		Page:    page,
		Offset:  offset,
		Filters: filters,
	}
}
