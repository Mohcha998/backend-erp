package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service": "inventory",
			"status":  "ok",
		})
	})

	log.Println("Inventory Service running on :8082")
	r.Run(":8082")
}
