package v1

import (
	"net/http"

	"auth-service/internal/domain"
	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	uc *usecase.MenuUsecase
}

func NewMenuHandler(uc *usecase.MenuUsecase) *MenuHandler {
	return &MenuHandler{uc}
}

func (h *MenuHandler) Create(c *gin.Context) {
	var menu domain.Menu
	c.ShouldBindJSON(&menu)

	if err := h.uc.Create(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, menu)
}

func (h *MenuHandler) GetAll(c *gin.Context) {
	data, _ := h.uc.GetAll()
	c.JSON(http.StatusOK, data)
}
