package routes

import (
	"github.com/Ngab-Rio/NOCs-API/internal/handlers"
	"github.com/Ngab-Rio/NOCs-API/internal/middleware"
	"github.com/Ngab-Rio/NOCs-API/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authHandler *handlers.AuthHandler, clientHandler *handlers.ClientHandler, jwtManager utils.JWTManager) {
	authRoutes := router.Group("/api/auth")
	AuthRoutes(authRoutes, authHandler)

	clientRoutes := router.Group("/api/clients")
	clientRoutes.Use(middleware.AuthMiddleware(jwtManager))
	ClientRoutes(clientRoutes, clientHandler)
}
