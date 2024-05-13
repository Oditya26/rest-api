package models

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageStr := c.DefaultQuery("page", "1")
		page, _ := strconv.Atoi(pageStr)
		if page <= 0 {
			page = 1
		}

		pageSizeStr := c.DefaultQuery("page_size", "5")
		pageSize, _ := strconv.Atoi(pageSizeStr)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		dbClone := db.Session((&gorm.Session{}))
		var total int64
		dbClone.Count(&total)

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

		c.Set("page", page)
		c.Set("pageSize", pageSize)
		c.Set("totalPages", totalPages)

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
