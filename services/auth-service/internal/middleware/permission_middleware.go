package middleware

import (
	"net/http"

	"auth-service/internal/repository"

	"github.com/gin-gonic/gin"
)

func RequirePermission(repo repository.UserRepository, permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")

		user, err := repo.FindWithRoles(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		for _, role := range user.Roles {
			for _, perm := range role.Permissions {
				if perm.Code == permission {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error": "permission denied",
		})
		c.Abort()
	}
}
