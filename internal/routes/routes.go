package routes

import (
	"github.com/Ngab-Rio/NOCs-API/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authHandler *handlers.AuthHandler) {
	authRoutes := router.Group("/api/auth")
	AuthRoutes(authRoutes, authHandler)
}
