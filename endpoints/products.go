package endpoints

import (
	"github.com/gin-gonic/gin"
	"main/handlers"
)

func RegisterProductEndpoint(router *gin.Engine) {
	products := router.Group("/products")
	{
		products.GET("/", handlers.GetProducts)
		products.GET("/:id", handlers.GetProduct)
		products.POST("/", handlers.StoreProduct)
		products.PATCH("/:id", handlers.UpdateProduct)
	}
}
