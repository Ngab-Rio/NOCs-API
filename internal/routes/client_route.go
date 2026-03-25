package routes

import (
	"github.com/Ngab-Rio/NOCs-API/internal/handlers"
	"github.com/gin-gonic/gin"
)

func ClientRoutes(r *gin.RouterGroup, clientHandler *handlers.ClientHandler) {
	r.POST("", clientHandler.CreateClient)
	r.PUT("/:id", clientHandler.UpdateClient)
	r.DELETE("/:id", clientHandler.DeleteClient)
	r.GET("/:id", clientHandler.GetClientByID)
}
