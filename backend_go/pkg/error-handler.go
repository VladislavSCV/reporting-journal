package pkg

import (
	"github.com/gin-gonic/gin"
)

// CError checks if the error is not nil and logs it. This is a convenience
// function for cases where you want to log an error but don't want to
// explicitly check if the error is not nil.
func CError(err error) {
	if err != nil {
		Log(err)
	}
}

func HandleHTTPError(c *gin.Context, statusCode int, message string, err error) {
	// Здесь можно добавить логику логирования ошибок, если это нужно
	CError(err) // Обрабатываем ошибку, например, логируем её
	c.JSON(statusCode, gin.H{
		"error":   message,
		"details": err.Error(),
	})
}
