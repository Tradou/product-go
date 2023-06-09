package endpoints

import "github.com/gin-gonic/gin"

func RegisterEndpoints() {
	router := gin.Default()

	RegisterProductEndpoint(router)
	RegisterAttributeEndpoint(router)
	RegisterAttributeProductEndpoint(router)
	RegisterOrderEndpoint(router)
	router.Run(":80")
}
