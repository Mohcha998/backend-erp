package middleware

import (
	"net/http"
	"strings"

	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

func AutoRefreshToken(authUC usecase.AuthUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.Next()
			return
		}

		// ✅ Validasi access token
		_, err := authUC.ValidateAccessToken(token)
		if err == nil {
			c.Next()
			return
		}

		// 🔁 Token expired → coba refresh
		refreshToken := c.GetHeader("X-Refresh-Token")
		if refreshToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "access token expired",
			})
			return
		}

		newAccess, newRefresh, err := authUC.RefreshToken(refreshToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "refresh token expired",
			})
			return
		}

		// Inject token baru
		c.Header("X-New-Access-Token", newAccess)
		c.Header("X-New-Refresh-Token", newRefresh)

		c.Next()
	}
}
