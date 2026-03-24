package handlers

import (
	"net/http"

	"github.com/Ngab-Rio/NOCs-API/internal/dto"
	"github.com/Ngab-Rio/NOCs-API/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body",
			"error":   err.Error(),
		})
		return
	}

	resp, err := h.authService.Login(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
		"data":    resp,
	})
}
