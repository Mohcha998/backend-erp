package v1

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(rg *gin.RouterGroup, h *AuthHandler) {
	v1 := rg.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", h.Login)
			auth.POST("/register", h.Register)
			auth.POST("/logout", h.Logout)
			auth.POST("/forgot-password", h.ForgotPassword)
		}
	}
}
