package role

import (
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
)

type RolePostgresRepository interface {
	CreateRole(*models.Role) error
	GetRoles() ([]*models.Role, error)
	GetRole(id int) (*models.Role, error)
	UpdateRole(*models.Role) error
	DeleteRole(id int) error
}

type RoleApiRepository interface {
	CreateRole(c *gin.Context) error
	GetRoles(c *gin.Context) error
	GetRole(c *gin.Context) error
	UpdateRole(c *gin.Context) error
	DeleteRole(c *gin.Context) error
}
