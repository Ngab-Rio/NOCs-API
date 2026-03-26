package utils

import (
	"net/http"

	appErr "github.com/Ngab-Rio/NOCs-API/internal/errors"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	switch err {
	case appErr.ErrNotFound:
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	case appErr.ErrInvalidRequest:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	case appErr.ErrUnauthorized:
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
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
