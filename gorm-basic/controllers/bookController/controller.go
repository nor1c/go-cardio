package bookcontroller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/nor1c/go-cardio/gorm-setup/models"
)

func GetBooks(c *gin.Context) {
	var books []models.BookResponse

	models.DB.Preload("Author").Find(&books)

	c.JSON(http.StatusOK, gin.H{"books": books})
}

func Create(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Request body invalid!", "err": err.Error()})
		return
	}

	if err := models.DB.Create(&book).Error; err != nil {
		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Duplicate on existing record!", "err": err.Error()})
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to publish your book, please try again later.", "err": err.Error()})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book published!", "book": book})
}
