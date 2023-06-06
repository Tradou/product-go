package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterProductEndpoint(router *gin.Engine) {
	users := router.Group("/products")
	{
		users.GET("/", getProducts)
	}
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Products",
	})
}
