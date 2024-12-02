package groups

import (
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
)

type GroupPostgresRepository interface {
	CreateGroup(group *models.Group) (int, error)
	GetGroupByID(id int) (*models.Group, error)
	GetGroups() ([]models.Group, error)
	UpdateGroup(group *models.Group) error
	DeleteGroup(id int) error

	GetCuratorGroups(id int) ([]models.Group, error)

	AddStudentToGroup(studentID, groupID int) error
	RemoveStudentFromGroup(studentID int) error
	GetStudentsByGroupID(groupID int) ([]*models.User, error)

	FindGroupsByName(name string) ([]*models.Group, error)
}

type GroupRedisRepository interface {
	CacheGroups(groups []*models.Group) error
	GetCachedGroups() ([]*models.Group, error)
}

type GroupApiRepository interface {
	CreateGroup(c *gin.Context)
	GetGroups(c *gin.Context)
	GetGroupByID(c *gin.Context)
	UpdateGroup(c *gin.Context)
	DeleteGroup(c *gin.Context)
}
