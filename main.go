package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func main() {
	router := gin.Default()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
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
