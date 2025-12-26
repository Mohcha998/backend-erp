package middleware

import (
	"auth-service/internal/pkg/apperror"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(*apperror.AppError); ok {
				c.JSON(appErr.HTTPStatus, appErr)
				return
			}

			c.JSON(500, apperror.ErrInternal)
		}
	}
}
