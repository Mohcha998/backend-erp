package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		limiter, exists := visitors[ip]
		if !exists {
			limiter = rate.NewLimiter(5, 10)
			visitors[ip] = limiter
		}
		mu.Unlock()

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			return
		}

		c.Next()
	}
}
