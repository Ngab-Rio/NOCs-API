package main

import (
	"time"

	"github.com/Ngab-Rio/NOCs-API/internal/config"
	"github.com/Ngab-Rio/NOCs-API/internal/database"
	"github.com/Ngab-Rio/NOCs-API/internal/handlers"
	"github.com/Ngab-Rio/NOCs-API/internal/repository"
	"github.com/Ngab-Rio/NOCs-API/internal/routes"
	"github.com/Ngab-Rio/NOCs-API/internal/services"
	"github.com/Ngab-Rio/NOCs-API/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cfg := config.Load()

	database.Connect(cfg.DBDSN())
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Is Running",
		})
	})

	jwtManager := utils.NewJWTManager(
		cfg.JWTSecret,
		cfg.JWTIssuer,
		1*time.Hour, // expired token
	)

	// AUTH
	authRepo := repository.NewAuthRepository(database.DB)
	authService := services.NewAuthService(authRepo, jwtManager)
	authHandler := handlers.NewAuthHandler(authService)

	// PACKAGE
	packageRepo := repository.NewPackageRepository(database.DB)
	packageService := services.NewPackageService(packageRepo)
	packageHandler := handlers.NewPackageHandler(packageService)

	routes.SetupRoutes(router, authHandler, packageHandler, jwtManager)

	router.Run(":" + cfg.AppPort)
}
