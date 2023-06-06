package handlers

import (
	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Store Product",
	})
}

func UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Product",
	})
}
