package groups

import (
	"github.com/VladislavSCV/internal/model"
)

type GroupPostgresRepository interface {
	CreateGroup(group *model.Group) error
	GetGroupByID(id int) (*model.Group, error)
	GetAllGroups() ([]*model.Group, error)
	UpdateGroup(group *model.Group) error
	DeleteGroup(id int) error

	AddStudentToGroup(studentID, groupID int) error
	RemoveStudentFromGroup(studentID int) error
	GetStudentsByGroupID(groupID int) ([]*model.User, error)

	FindGroupsByName(name string) ([]*model.Group, error)
}

type GroupRedisRepository interface {
	CacheGroups(groups []*model.Group) error
	GetCachedGroups() ([]*model.Group, error)
}
