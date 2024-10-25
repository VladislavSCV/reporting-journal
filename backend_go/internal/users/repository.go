package users

import (
	"github.com/VladislavSCV/internal/model"
	"github.com/gin-gonic/gin"
)

type Schedule struct {
	ID         int    `json:"id"`
	GroupID    int    `json:"group_id"`
	DayOfWeek  int    `json:"day_of_week"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Subject    string `json:"subject"`
	TeacherID  int    `json:"teacher_id"`
	Location   string `json:"location"`
	Recurrence string `json:"recurrence"`
}

type UserPostgresRepository interface {
	GetUsers() (*[]model.User, error)
	GetUserByLogin(login string) (model.User, error)
	GetUserById(id int) (model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(id string, updates map[string]string) error
	DeleteUser(id int) error
}

type UserRedisRepository interface {
	Login(user *model.User) error
	Logout(id int) error
	GetUserById(id int) (model.User, error)
	UpdateUser(id string, updates map[string]string) error
	DeleteUser(id int) error
}

type UserAPIRepository interface {
	Login(c *gin.Context) (model.User, error)
	SignUp(c *gin.Context) error
	GetUser(c *gin.Context) error
	GetUserByLogin(c *gin.Context) (model.User, error)
	GetUsers(c *gin.Context) error
	UpdateUser(c *gin.Context) error
	DeleteUser(c *gin.Context) error
}
