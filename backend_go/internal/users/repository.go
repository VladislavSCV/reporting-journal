package users

import (
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
)

type Schedule struct {
	ID         int    `json:"id"`
	GroupID    int    `json:"group_id"`
	DayOfWeek  int    `json:"day_of_week"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Subject    string `json:"subjects"`
	TeacherID  int    `json:"teacher_id"`
	Location   string `json:"location"`
	Recurrence string `json:"recurrence"`
}

type UserPostgresRepository interface {
	GetUsers() ([]models.User, error)
	GetStudents() ([]models.User, error)
	GetTeachers() ([]models.User, error)
	GetUsersByGroupID(groupID int) ([]models.User, error)
	GetUsersByRoleID(roleID int) ([]models.User, error)
	GetUserByToken(token string) (models.User, error)
	GetUserByLogin(login string) (models.User, error)
	GetUserById(id int) (models.User, error)
	CreateStudent(user *models.User) (models.User, string, error)
	CreateTeacher(user *models.User) (models.User, string, error)
	UpdateUser(id int, updates map[string]string) error
	UpdateToken(id int, token string) error
	DeleteUser(id int) error
}

type UserRedisRepository interface {
	SaveInCache(user *models.User) error
	Logout(id int) error
	UpdateUser(id string, updates map[string]string) error
	DeleteUser(id int) error
}

type UserAPIRepository interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
	VerifyToken(c *gin.Context)
	GetUser(c *gin.Context)
	GetUserByLogin(c *gin.Context)
	GetUserByToken(c *gin.Context)
	GetUsers(c *gin.Context)
	GetStudents(c *gin.Context)
	GetTeachers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
