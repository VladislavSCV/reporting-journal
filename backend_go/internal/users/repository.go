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
	GetUserByToken(token string) (string, error)
	GetUserByLogin(login string) (models.User, error)
	GetUserById(id int) (models.User, error)
	CreateStudent(user *models.User) (string, error)
	CreateTeacher(user *models.User) (string, error)
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
	Login(c *gin.Context) error
	SignUp(c *gin.Context) error
	VerifyToken(c *gin.Context) error
	GetUser(c *gin.Context) error
	GetUserByLogin(c *gin.Context) (models.User, error)
	GetUserByToken(c *gin.Context) error
	GetUsers(c *gin.Context) error
	GetStudents(c *gin.Context) error
	GetTeachers(c *gin.Context) error
	UpdateUser(c *gin.Context) error
	DeleteUser(c *gin.Context) error
}
