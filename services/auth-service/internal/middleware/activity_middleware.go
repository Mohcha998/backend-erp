package middleware

import (
	"auth-service/internal/domain"
	"auth-service/internal/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func ActivityLog(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		db := database.GetDB()
		if db == nil {
			return
		}

		userID, ok := c.Get("user_id")
		if !ok {
			return
		}

		db.Create(&domain.ActivityLog{
			UserID:    userID.(uint),
			Activity:  action,
			IPAddress: c.ClientIP(),
		})
	}
}
