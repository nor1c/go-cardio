package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nor1c/GoGinGorm/models"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProduct(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Request body not accepted!"})
		return
	}

	models.DB.Create(&product)

	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Request body not accepted!"})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotModified, gin.H{"message": "Failed to update product!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated!"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusExpectationFailed, gin.H{"message": "Failed to delete product!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product removed!"})
}
