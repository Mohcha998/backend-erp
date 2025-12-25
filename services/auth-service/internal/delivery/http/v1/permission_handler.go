package v1

import (
	"net/http"

	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RolePermissionHandler struct {
	uc *usecase.PermissionUsecase
}

func NewRolePermissionHandler(uc *usecase.PermissionUsecase) *RolePermissionHandler {
	return &RolePermissionHandler{uc}
}

/*
Request JSON:

	{
	  "role_id": 1,
	  "permission_ids": [1,2,3]
	}
*/
func (h *RolePermissionHandler) AssignPermission(c *gin.Context) {
	var req struct {
		RoleID        uint   `json:"role_id"`
		PermissionIDs []uint `json:"permission_ids"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Assign(req.RoleID, req.PermissionIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "permission assigned to role successfully",
	})
}
