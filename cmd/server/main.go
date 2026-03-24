package main

import (
	"github.com/Ngab-Rio/NOCs-API/internal/config"
	"github.com/Ngab-Rio/NOCs-API/internal/database"
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

	router.Run(":" + cfg.AppPort)
}
