package v1

import (
	"net/http"

	"auth-service/internal/domain"
	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	uc *usecase.PermissionUsecase
}

func NewPermissionHandler(uc *usecase.PermissionUsecase) *PermissionHandler {
	return &PermissionHandler{uc}
}

func (h *PermissionHandler) Create(c *gin.Context) {
	var p domain.RoleMenuPermission
	c.ShouldBindJSON(&p)

	if err := h.uc.Create(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (h *PermissionHandler) GetAll(c *gin.Context) {
	data, _ := h.uc.GetAll()
	c.JSON(http.StatusOK, data)
}
