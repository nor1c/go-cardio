package authorcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nor1c/go-cardio/gorm-setup/models"
)

func GetAuthors(c *gin.Context) {
	var authors []models.Author

	models.DB.Preload("Books").Find(&authors)

	c.JSON(http.StatusOK, gin.H{"authors": &authors})
}

func Create(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Request body invalid!", "err": err.Error()})
		return
	}

	models.DB.Create(&author)

	c.JSON(http.StatusCreated, gin.H{"message": "Author registered!", "author": author})
}
