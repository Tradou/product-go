package handlers

import (
	"github.com/gin-gonic/gin"
	"main/infrastructure/models"
	"net/http"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Products",
	})
}

func GetProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Product",
	})
}

func StoreProduct(c *gin.Context) {
	var product models.StoreProduct
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Product",
	})
}
