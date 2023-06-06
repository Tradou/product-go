package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterProductEndpoint(router *gin.Engine) {
	products := router.Group("/products")
	{
		products.GET("/", getProducts)
	}
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Products",
	})
}
