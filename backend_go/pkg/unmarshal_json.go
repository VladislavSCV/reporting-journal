package pkg

import (
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
)

// ParseUserRequest UnmarshalJson reads the JSON request body from gin.Context
func ParseUserRequest(c *gin.Context) (*models.User, error) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
