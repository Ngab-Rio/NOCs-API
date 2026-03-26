package handlers

import (
	"strconv"

	"github.com/Ngab-Rio/NOCs-API/internal/dto"
	"github.com/Ngab-Rio/NOCs-API/internal/services"
	"github.com/Ngab-Rio/NOCs-API/internal/utils"
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
		utils.HandleError(c, err)
		return
	}

	resp, err := h.packageService.CreatePackage(c, req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Package Created", resp)
}

func (h *PackageHandler) UpdatePackage(c *gin.Context) {
	packageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var req dto.UpdatePackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleError(c, err)
		return
	}

	resp, err := h.packageService.UpdatePackage(c, packageID, &req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Package Updated", resp)
}

func (h *PackageHandler) DeletePackage(c *gin.Context) {
	packageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	err = h.packageService.DeletePackage(c, packageID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Package Deleted", nil)
}

func (h *PackageHandler) FindPackageByID(c *gin.Context) {
	packageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	pkg, err := h.packageService.FindPackageByID(c, packageID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Package Found", pkg)
}

func (h *PackageHandler) FindAllPackages(c *gin.Context) {
	packages, err := h.packageService.FindAllPackages(c)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Packages Found", packages)
}
