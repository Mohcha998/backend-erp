package v1

import (
	"net/http"
	"strconv"

	"auth-service/internal/domain"
	"auth-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

// ////////////////////////////////////////////////////
// CREATE USER
// ////////////////////////////////////////////////////

// Create godoc
// @Summary Create new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body domain.User true "User payload"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		DivisionID uint   `json:"division_id"`
		IsActive   bool   `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := &domain.User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   string(hash),
		DivisionID: req.DivisionID,
		IsActive:   true,
	}

	if err := h.uc.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})
}

// ////////////////////////////////////////////////////
// GET ALL USERS
// ////////////////////////////////////////////////////

// GetAll godoc
// @Summary Get all users
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {array} domain.User
// @Router /v1/users [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// ////////////////////////////////////////////////////
// GET USER BY ID
// ////////////////////////////////////////////////////

// GetByID godoc
// @Summary Get user by ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} domain.User
// @Failure 404 {object} map[string]string
// @Router /v1/users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ////////////////////////////////////////////////////
// UPDATE USER
// ////////////////////////////////////////////////////

// Update godoc
// @Summary Update user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param body body domain.User true "User payload"
// @Success 200 {object} map[string]string
// @Router /v1/users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Password   string `json:"password,omitempty"`
		DivisionID uint   `json:"division_id"`
		IsActive   bool   `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &domain.User{
		ID:         uint(id),
		Name:       req.Name,
		Email:      req.Email,
		DivisionID: req.DivisionID,
		IsActive:   req.IsActive,
	}

	if req.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			return
		}
		user.Password = string(hash)
	}

	if err := h.uc.Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
	})
}

// ////////////////////////////////////////////////////
// DELETE USER
// ////////////////////////////////////////////////////

// Delete godoc
// @Summary Delete user
// @Tags Users
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Router /v1/users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}
