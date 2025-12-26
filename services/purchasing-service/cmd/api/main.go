package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service": "purchasing",
			"status":  "ok",
		})
	})

	log.Println("Purchasing Service running on :8083")
	r.Run(":8083")
}
