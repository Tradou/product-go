package endpoints

import (
	"github.com/gin-gonic/gin"
	"main/handlers"
)

func RegisterAttributeEndpoint(router *gin.Engine) {
	products := router.Group("/attributes")
	{
		products.GET("", handlers.GetAttributes)
		products.GET("/:id", handlers.GetAttribute)
		products.POST("", handlers.StoreAttribute)
		products.PATCH("/:id", handlers.UpdateAttribute)
	}
}
