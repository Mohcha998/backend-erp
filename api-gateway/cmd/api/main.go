package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/middleware"
	"api-gateway/internal/proxy"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	r := gin.New()

	r.Use(
		gin.Recovery(),
		middleware.RequestID(),
		middleware.CORS(),
		middleware.RateLimiter(),
		middleware.JWTPrecheck(),
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	api := r.Group("/api")

	api.Any("/auth/*path", proxy.Forward(cfg.AuthServiceURL))
	api.Any("/inventory/*path", proxy.Forward(cfg.InventoryServiceURL))
	api.Any("/purchasing/*path", proxy.Forward(cfg.PurchasingServiceURL))

	r.Run(":" + cfg.Port)
}

//test
