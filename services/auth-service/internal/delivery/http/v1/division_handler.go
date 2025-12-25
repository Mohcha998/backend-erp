package v1

import (
	"net/http"
	"strconv"

	"auth-service/internal/domain"
	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type DivisionHandler struct{ uc *usecase.DivisionUsecase }

func NewDivisionHandler(u *usecase.DivisionUsecase) *DivisionHandler {
	return &DivisionHandler{u}
}

func (h *DivisionHandler) Create(c *gin.Context) {
	var d domain.Division
	c.ShouldBindJSON(&d)
	h.uc.Create(&d)
	c.JSON(http.StatusCreated, d)
}

func (h *DivisionHandler) GetAll(c *gin.Context) {
	data, _ := h.uc.GetAll()
	c.JSON(http.StatusOK, data)
}

func (h *DivisionHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.uc.Delete(uint(id))
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
