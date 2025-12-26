package v1

import (
	"net/http"
	"strconv"

	"auth-service/internal/domain"
	"auth-service/internal/usecase"
	"auth-service/internal/pkg/apperror"

	"github.com/gin-gonic/gin"
)

type DivisionHandler struct {
	uc *usecase.DivisionUsecase
}

func NewDivisionHandler(u *usecase.DivisionUsecase) *DivisionHandler {
	return &DivisionHandler{uc: u}
}

// CreateDivision godoc
// @Summary Create division
// @Description Create new division
// @Tags Division
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body domain.Division true "Division payload"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/divisions [post]
func (h *DivisionHandler) Create(c *gin.Context) {
	var req domain.Division

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.ErrBadRequest.Message})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "division name is required"})
		return
	}

	if err := h.uc.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Message})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "division created successfully",
	})
}

// GetAllDivisions godoc
// @Summary Get all divisions
// @Description Get list of divisions
// @Tags Division
// @Produce json
// @Security BearerAuth
// @Success 200 {array} domain.Division
// @Failure 500 {object} map[string]string
// @Router /v1/divisions [get]
func (h *DivisionHandler) GetAll(c *gin.Context) {
	data, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Message})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetDivisionByID godoc
// @Summary Get division by ID
// @Description Get detail division
// @Tags Division
// @Produce json
// @Security BearerAuth
// @Param id path int true "Division ID"
// @Success 200 {object} domain.Division
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /v1/divisions/{id} [get]
func (h *DivisionHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	data, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": apperror.ErrNotFound.Message})
		return
	}

	c.JSON(http.StatusOK, data)
}

// UpdateDivision godoc
// @Summary Update division
// @Description Update division data
// @Tags Division
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Division ID"
// @Param body body domain.Division true "Division payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/divisions/{id} [put]
func (h *DivisionHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req domain.Division
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = uint(id)

	if err := h.uc.Update(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "division updated successfully",
	})
}

// DeleteDivision godoc
// @Summary Delete division
// @Description Soft delete division
// @Tags Division
// @Produce json
// @Security BearerAuth
// @Param id path int true "Division ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/divisions/{id} [delete]
func (h *DivisionHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": apperror.ErrInternal.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "division deleted successfully",
	})
}
