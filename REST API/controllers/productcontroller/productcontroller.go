package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tentangcode/go-restapi-gin/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var bukus []models.Buku

	models.DB.Order("judul asc").Find(&bukus)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data ditemukan",
		"data":    bukus,
	})
}

func Show(c *gin.Context) {
	var buku models.Buku
	id := c.Param("id")

	if err := models.DB.First(&buku, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "Data tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": err.Error(),
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data ditemukan",
		"data":    buku,
	})
}

func Create(c *gin.Context) {
	var buku models.Buku

	if err := c.ShouldBindJSON(&buku); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Gagal memasukkan data",
		})
		return
	}
	validate := validator.New()
	if err := validate.Struct(buku); err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": errors.Error(),
		})
		return
	}

	if err := models.DB.Create(&buku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Gagal memasukkan data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Sukses memasukkan data",
	})
}

func Update(c *gin.Context) {
	var buku models.Buku
	id := c.Param("id")

	if err := c.ShouldBindJSON(&buku); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Gagal melakukan update data",
		})
		return
	}

	if models.DB.Model(&buku).Where("id = ?", id).Updates(&buku).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Data tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Sukses melakukan update data",
	})
}

func Delete(c *gin.Context) {
	var buku models.Buku

	id := c.Param("id")

	if models.DB.Delete(&buku, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Sukses melakukan delete data",
	})
}
