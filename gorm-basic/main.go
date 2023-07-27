package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nor1c/go-cardio/gorm-setup/models"

	authorController "github.com/nor1c/go-cardio/gorm-setup/controllers/authorController"
	bookController "github.com/nor1c/go-cardio/gorm-setup/controllers/bookController"
)

func main() {
	models.Connect()

	// gin router
	r := gin.Default()

	authorR := r.Group("/authors")
	{
		authorR.GET("/", authorController.GetAuthors)
		authorR.POST("/", authorController.Create)
	}

	bookR := r.Group("/books")
	{
		bookR.GET("/", bookController.GetBooks)
		bookR.POST("/", bookController.Create)
	}

	r.Run(":8080")
}
