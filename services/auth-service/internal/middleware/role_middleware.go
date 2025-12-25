package middleware

import (
	"net/http"

	"auth-service/internal/repository"

	"github.com/gin-gonic/gin"
)

func RequireRole(repo repository.UserRepository, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")

		user, err := repo.FindWithRoles(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		for _, role := range user.Roles {
			for _, allowed := range allowedRoles {
				if role.Name == allowed {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error": "you do not have permission",
		})
		c.Abort()
	}
}
