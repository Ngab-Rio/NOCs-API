package handlers

import (
	"net/http"
	"strconv"

	"github.com/Ngab-Rio/NOCs-API/internal/dto"
	"github.com/Ngab-Rio/NOCs-API/internal/services"
	"github.com/gin-gonic/gin"
)

type PackageHandler struct {
	packageService services.PackageService
}

func NewPackageHandler(packageService services.PackageService) *PackageHandler {
	return &PackageHandler{packageService: packageService}
}

func (h *PackageHandler) CreatePackage(c *gin.Context) {
	var req dto.CreatePackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body",
			"error":   err.Error(),
		})
		return
	}

	resp, err := h.packageService.CreatePackage(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Package Created",
		"data":    resp,
	})
}

func (h *PackageHandler) UpdatePackage(c *gin.Context) {
	packageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Package ID",
			"error":   err.Error(),
		})
		return
	}

	var req dto.UpdatePackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request Body",
			"error":   err.Error(),
		})
		return
	}

	resp, err := h.packageService.UpdatePackage(c, packageID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Package Updated",
		"data":    resp,
	})
}

func (h *PackageHandler) DeletePackage(c *gin.Context) {
	packageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Package ID",
			"error":   err.Error(),
		})
		return
	}

	err = h.packageService.DeletePackage(c, packageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Package Deleted",
	})
}

func (h *PackageHandler) FindPackageByID(c *gin.Context) {
	packageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Package ID",
			"error":   err.Error(),
		})
		return
	}

	pkg, err := h.packageService.FindPackageByID(c, packageID)
	if err != nil {
		if err.Error() == "package not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Package not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Package Found",
		"data":    pkg,
	})
}

func (h *PackageHandler) FindAllPackages(c *gin.Context) {
	packages, err := h.packageService.FindAllPackages(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Packages Found",
		"data":    packages,
	})
}
