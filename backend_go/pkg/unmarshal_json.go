package pkg

import (
	"github.com/VladislavSCV/internal/model"
	"github.com/gin-gonic/gin"
)

// ParseUserRequest UnmarshalJson reads the JSON request body from gin.Context
func ParseUserRequest(c *gin.Context) (*model.User, error) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
