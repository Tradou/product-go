package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
	"main/infrastructure/migrations"
	"net/http"
	"os"
)

func main() {
	migrations.Migrate()

	router := gin.Default()

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		),
		DB: 0,
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ssss",
		})
	})

	redisRoute := router.Group("/redis")
	{
		redisRoute.GET("/write", func(c *gin.Context) {
			err := rdb.LPush(context.Background(), "dummy", "hahahahahaha").Err()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Data written successfully"})
		})

		redisRoute.GET("/retrieve", func(c *gin.Context) {
			result, err := rdb.LRange(context.Background(), "dummy", 0, -1).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from Redis"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": result})
		})
	}

	router.Run(":80")
}
