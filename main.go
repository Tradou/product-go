package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"main/endpoints"
	"main/infrastructure/migrations"
)

func main() {
	migrations.Migrate()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	endpoints.RegisterProductEndpoint(router)

	router.Run(":80")
}
