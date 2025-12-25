package v1

import (
	"net/http"

	"auth-service/internal/domain"
	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	uc *usecase.RoleUsecase
}

func NewRoleHandler(uc *usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{uc}
}

func (h *RoleHandler) Create(c *gin.Context) {
	var role domain.Role
	c.ShouldBindJSON(&role)

	if err := h.uc.Create(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

func (h *RoleHandler) GetAll(c *gin.Context) {
	data, _ := h.uc.GetAll()
	c.JSON(http.StatusOK, data)
}
