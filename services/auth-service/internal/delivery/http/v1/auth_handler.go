package v1

import (
	"net/http"
	"strings"

	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC usecase.AuthUsecase
}

func NewAuthHandler(authUC usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC}
}

// Login godoc
// @Summary Login
// @Description Login user and return access & refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body map[string]string true "email & password"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.authUC.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"token_type":    "Bearer",
	})
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Generate new access token using refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body map[string]string true "refresh_token"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.authUC.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Register godoc
// @Summary Register user
// @Description Register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "user payload"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Name       string `json:"name" binding:"required"`
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required,min=6"`
		DivisionID uint   `json:"division_id" binding:"required"`
		RoleID     uint   `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.authUC.Register(
		req.Name,
		req.Email,
		req.Password,
		req.DivisionID,
		req.RoleID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
	})
}

// Logout godoc
// @Summary Logout user
// @Description Revoke refresh token
// @Tags Auth
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		return
	}

	h.authUC.Logout(token)

	c.JSON(http.StatusOK, gin.H{
		"message": "logged out successfully",
	})
}

// ForgotPassword godoc
// @Summary Forgot password
// @Description Generate reset token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body map[string]string true "email"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUC.ForgotPassword(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ⚠️ Production: kirim via email
	c.JSON(http.StatusOK, gin.H{
		"reset_token": token,
	})
}
