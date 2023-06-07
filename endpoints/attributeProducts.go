package endpoints

import (
	"github.com/gin-gonic/gin"
	"main/handlers"
)

func RegisterAttributeProductEndpoint(router *gin.Engine) {
	products := router.Group("/attribute_products")
	{
		products.GET("", handlers.GetAttributeProducts)
		products.GET("/:id", handlers.GetAttributeProduct)
		products.POST("", handlers.StoreAttributeProduct)
		products.PATCH("/:id", handlers.UpdateAttributeProduct)
	}
}
