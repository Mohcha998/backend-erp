package v1

import (
	"net/http"
	"strconv"

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

// =====================================================
// CREATE ROLE
// =====================================================

// CreateRole godoc
// @Summary Create role
// @Description Create new role
// @Tags Roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body domain.Role true "Role data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/roles [post]
func (h *RoleHandler) Create(c *gin.Context) {
	var req domain.Role

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role name is required"})
		return
	}

	if err := h.uc.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "role created successfully",
	})
}

// =====================================================
// GET ALL ROLES
// =====================================================

// GetRoles godoc
// @Summary Get all roles
// @Tags Roles
// @Produce json
// @Security BearerAuth
// @Success 200 {array} domain.Role
// @Failure 500 {object} map[string]string
// @Router /v1/roles [get]
func (h *RoleHandler) GetAll(c *gin.Context) {
	data, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// =====================================================
// GET ROLE BY ID
// =====================================================

// GetRoleByID godoc
// @Summary Get role by ID
// @Tags Roles
// @Produce json
// @Security BearerAuth
// @Param id path int true "Role ID"
// @Success 200 {object} domain.Role
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /v1/roles/{id} [get]
func (h *RoleHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}

	data, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}

// =====================================================
// UPDATE ROLE
// =====================================================

// UpdateRole godoc
// @Summary Update role
// @Tags Roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Role ID"
// @Param body body domain.Role true "Role data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/roles/{id} [put]
func (h *RoleHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}

	var req domain.Role
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = uint(id)

	if err := h.uc.Update(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "role updated successfully",
	})
}

// =====================================================
// DELETE ROLE
// =====================================================

// DeleteRole godoc
// @Summary Delete role
// @Tags Roles
// @Security BearerAuth
// @Param id path int true "Role ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/roles/{id} [delete]
func (h *RoleHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}

	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "role deleted successfully",
	})
}
