package v1

import (
	"net/http"
	"strconv"

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

// CreateMenu godoc
// @Summary Create new menu
// @Description Create new menu data
// @Tags Menu
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body domain.Menu true "Menu payload"
// @Success 201 {object} domain.Menu
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/menus [post]
func (h *MenuHandler) Create(c *gin.Context) {
	var req domain.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

// GetAllMenus godoc
// @Summary Get all menus
// @Description Retrieve all menu data
// @Tags Menu
// @Produce json
// @Security BearerAuth
// @Success 200 {array} domain.Menu
// @Failure 500 {object} map[string]string
// @Router /v1/menus [get]
func (h *MenuHandler) GetAll(c *gin.Context) {
	data, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetMenuByID godoc
// @Summary Get menu by ID
// @Description Get menu detail by ID
// @Tags Menu
// @Produce json
// @Security BearerAuth
// @Param id path int true "Menu ID"
// @Success 200 {object} domain.Menu
// @Failure 404 {object} map[string]string
// @Router /v1/menus/{id} [get]
func (h *MenuHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "menu not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}

// UpdateMenu godoc
// @Summary Update menu
// @Description Update menu data
// @Tags Menu
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Menu ID"
// @Param body body domain.Menu true "Menu payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/menus/{id} [put]
func (h *MenuHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req domain.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = uint(id)

	if err := h.uc.Update(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "menu updated"})
}

// DeleteMenu godoc
// @Summary Delete menu
// @Description Delete menu by ID
// @Tags Menu
// @Produce json
// @Security BearerAuth
// @Param id path int true "Menu ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/menus/{id} [delete]
func (h *MenuHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "menu deleted"})
}
