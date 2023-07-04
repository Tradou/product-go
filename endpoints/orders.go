package endpoints

import (
	"github.com/gin-gonic/gin"
	"main/handlers"
)

func RegisterOrderEndpoint(router *gin.Engine) {
	orders := router.Group("/orders")
	{
		orders.GET("", handlers.GetOrders)
		orders.GET("/:id", handlers.GetOrder)
	}
}
