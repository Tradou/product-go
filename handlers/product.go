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
