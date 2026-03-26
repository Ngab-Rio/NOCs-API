package handlers

import (
	"github.com/Ngab-Rio/NOCs-API/internal/dto"
	"github.com/Ngab-Rio/NOCs-API/internal/services"
	"github.com/Ngab-Rio/NOCs-API/internal/utils"
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
		utils.HandleError(c, err)
		return
	}

	resp, err := h.authService.Login(c, req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Login Success", resp)
}
