package v1

import (
	"net/http"

	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RoleMenuHandler struct {
	uc *usecase.RoleMenuUsecase
}

func NewRoleMenuHandler(uc *usecase.RoleMenuUsecase) *RoleMenuHandler {
	return &RoleMenuHandler{uc}
}

/*
Request JSON:

	{
	  "role_id": 1,
	  "menu_ids": [1,2,3]
	}
*/
func (h *RoleMenuHandler) AssignMenu(c *gin.Context) {
	var req struct {
		RoleID  uint   `json:"role_id"`
		MenuIDs []uint `json:"menu_ids"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Assign(req.RoleID, req.MenuIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "menu assigned to role successfully",
	})
}
