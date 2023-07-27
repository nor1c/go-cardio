package main

import (
	"github.com/gin-gonic/gin"
	productController "github.com/nor1c/GoGinGorm/controllers/productController"
	"github.com/nor1c/GoGinGorm/models"
)

func main() {
	r := gin.Default()
	models.Connect()

	r.GET("/products", productController.GetAll)
	r.GET("/product/:id", productController.GetProduct)
	r.POST("/product", productController.Create)
	r.PUT("/product/:id", productController.Update)
	r.DELETE("/product/:id", productController.Delete)

	r.Run()
}
