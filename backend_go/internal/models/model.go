package models

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

const (
	MONDAY    = "Monday"
	TUESDAY   = "Tuesday"
	WEDNESDAY = "Wednesday"
	THURSDAY  = "Thursday"
	FRIDAY    = "Friday"
	SATURDAY  = "Saturday"
	SUNDAY    = "Sunday"
)

type UserAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	RoleID     int    `json:"role_id"`
	GroupID    *int   `json:"group_id"`
	Login      string `json:"login"`
	Hash       string `json:"password"`
	Salt       string `json:"salt"`
	Token      string `json:"token"`

	Status *string `json:"status"`
	Role   string  `json:"role"`
	Group  *string `json:"group"`
}

type Role struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	GroupId *int   `json:"group_id"`
	UserId  int    `json:"user_id"`
}

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Schedule struct {
	ScheduleID  int
	GroupName   string
	DayOfWeek   int
	SubjectName string
	TeacherName string
	Location    string
}

type CreateSchedule struct {
	ID        int    `json:"id"`
	GroupID   int    `json:"group_id"`
	DayOfWeek int    `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	SubjectID int    `json:"subject_id"`
	TeacherID int    `json:"teacher_id"`
	Location  string `json:"location"`
}

type Subject struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Execer интерфейс для работы с транзакцией или базой данных.
type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type Else interface {
	GetAdminPanelData(c *gin.Context)
	GetCuratorGroupsStudentList(c *gin.Context)
	StudentsAttendance(c *gin.Context)
	UpdateAttendance(c *gin.Context)
}
