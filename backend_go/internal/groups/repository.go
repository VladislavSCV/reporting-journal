package groups

import (
	"github.com/VladislavSCV/internal/models"
)

type GroupPostgresRepository interface {
	CreateGroup(group *models.Group) error
	GetGroupByID(id int) (*models.Group, error)
	GetAllGroups() ([]*models.Group, error)
	UpdateGroup(group *models.Group) error
	DeleteGroup(id int) error

	AddStudentToGroup(studentID, groupID int) error
	RemoveStudentFromGroup(studentID int) error
	GetStudentsByGroupID(groupID int) ([]*models.User, error)

	FindGroupsByName(name string) ([]*models.Group, error)
}

type GroupRedisRepository interface {
	CacheGroups(groups []*models.Group) error
	GetCachedGroups() ([]*models.Group, error)
}
