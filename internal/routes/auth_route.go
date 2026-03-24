package routes

import (
	"github.com/Ngab-Rio/NOCs-API/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	r.POST("/login", authHandler.Login)
}
