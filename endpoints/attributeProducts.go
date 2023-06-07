package endpoints

import (
	"github.com/gin-gonic/gin"
	"main/handlers"
)

func RegisterAttributeProductEndpoint(router *gin.Engine) {
	products := router.Group("/attributeproducts")
	{
		products.GET("/", handlers.GetAttributeProducts)
		products.GET("/:id", handlers.GetAttributeProduct)
		products.POST("/", handlers.StoreAttributeProduct)
		products.PATCH("/:id", handlers.UpdateAttributeProduct)
	}
}
