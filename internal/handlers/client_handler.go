package handlers

import (
	"strconv"

	"github.com/Ngab-Rio/NOCs-API/internal/dto"
	"github.com/Ngab-Rio/NOCs-API/internal/services"
	"github.com/Ngab-Rio/NOCs-API/internal/utils"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	clientService services.ClientService
}

func NewClientHandler(clientService services.ClientService) *ClientHandler {
	return &ClientHandler{clientService: clientService}
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var req dto.CreateClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleError(c, err)
		return
	}

	resp, err := h.clientService.CreateClient(c, req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Client Created Successfully", resp)
}

func (h *ClientHandler) UpdateClient(c *gin.Context) {
	var req dto.UpdateClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleError(c, err)
		return
	}

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp, err := h.clientService.UpdateClient(c, int(id), req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Client Updated Successfully", resp)
}

func (h *ClientHandler) DeleteClient(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	err = h.clientService.DeleteClient(c, int(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Client Deleted Successfully", nil)
}

func (h *ClientHandler) GetClientByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp, err := h.clientService.GetClientByID(c, int(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Client Found", resp)
}

func (h *ClientHandler) GetClients(c *gin.Context) {
	var req dto.GetClientsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.HandleError(c, err)
		return
	}

	resp, err := h.clientService.GetClients(c, req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Clients Found", resp)
}
