package endpoints

import "github.com/gin-gonic/gin"

func RegisterEndpoints() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	RegisterProductEndpoint(router)
	router.Run(":80")
}
