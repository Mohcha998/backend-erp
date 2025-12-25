package middleware

import (
	"net/http"

	"auth-service/internal/repository"

	"github.com/gin-gonic/gin"
)

func IsActiveUser(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")

		user, err := userRepo.FindByID(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		if !user.IsActive {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "your account is inactive",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
