package routes

import (
	"github.com/Ngab-Rio/NOCs-API/internal/handlers"
	"github.com/Ngab-Rio/NOCs-API/internal/middleware"
	"github.com/Ngab-Rio/NOCs-API/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authHandler *handlers.AuthHandler, packageHandler *handlers.PackageHandler, jwtManager utils.JWTManager) {
	authRoutes := router.Group("/api/auth")
	AuthRoutes(authRoutes, authHandler)

	packageRoutes := router.Group("/api/package")
	packageRoutes.Use(middleware.AuthMiddleware(jwtManager))
	PackageRoutes(packageRoutes, packageHandler)
}
