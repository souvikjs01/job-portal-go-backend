package handlers

import (
	"job_portal/packages/models"
	"job_portal/packages/services"
	"job_portal/packages/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req models.CreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   "Invalid request payload",
		})
		return
	}

	if err := validation.ValidateStruct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	// Call service layer
	user, token, err := h.userService.Register(&req)
	if err != nil || token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
		"data":    user,
	})
}
