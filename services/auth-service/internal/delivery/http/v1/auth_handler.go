package v1

import (
	"net/http"
	"strings"

	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authUC usecase.AuthUsecase
}

func NewAuthHandler(authUC usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC}
}

/* ========== LOGIN ========== */
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUC.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

/* ========== REGISTER ========== */
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		DivisionID string `json:"division_id"`
		RoleID     string `json:"role_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	divID, err := uuid.Parse(req.DivisionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid division_id"})
		return
	}

	roleID, err := uuid.Parse(req.RoleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role_id"})
		return
	}

	err = h.authUC.Register(req.Name, req.Email, req.Password, divID, roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

/* ========== LOGOUT ========== */
func (h *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := strings.Replace(authHeader, "Bearer ", "", 1)

	h.authUC.Logout(token)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

/* ========== FORGOT PASSWORD ========== */
func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
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

	// nanti token ini dikirim via email
	c.JSON(http.StatusOK, gin.H{"reset_token": token})
}
