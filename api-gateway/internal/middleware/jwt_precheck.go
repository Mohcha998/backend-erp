package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTPrecheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.Next()
			return
		}

		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token format",
			})
			return
		}

		c.Next()
	}
}
