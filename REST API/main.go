package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tentangcode/go-restapi-gin/controllers/productcontroller"
	"github.com/tentangcode/go-restapi-gin/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/buku", productcontroller.Index)
	r.GET("/api/buku/:id", productcontroller.Show)
	r.POST("/api/buku", productcontroller.Create)
	r.PUT("/api/buku/:id", productcontroller.Update)
	r.DELETE("/api/buku/:id", productcontroller.Delete)

	r.Run()
}
