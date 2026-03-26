package utils

import (
	"errors"
	"net/http"

	appErr "github.com/Ngab-Rio/NOCs-API/internal/errors"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, appErr.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return

	case errors.Is(err, appErr.ErrInvalidRequest):
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return

	case errors.Is(err, appErr.ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return

	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}
