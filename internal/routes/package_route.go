package routes

import (
	"github.com/Ngab-Rio/NOCs-API/internal/handlers"
	"github.com/gin-gonic/gin"
)

func PackageRoutes(router *gin.RouterGroup, packageHandler *handlers.PackageHandler) {
	router.POST("", packageHandler.CreatePackage)
	router.PUT("/:id", packageHandler.UpdatePackage)
	router.DELETE("/:id", packageHandler.DeletePackage)
	router.GET("/:id", packageHandler.FindPackageByID)
	router.GET("", packageHandler.FindAllPackages)
}
