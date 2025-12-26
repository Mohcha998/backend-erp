package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if limiter, ok := visitors[ip]; ok {
		return limiter
	}

	limiter := rate.NewLimiter(5, 10)
	visitors[ip] = limiter
	return limiter
}

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter := getLimiter(c.ClientIP())
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}
		c.Next()
	}
}
