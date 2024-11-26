package schedules

import (
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
)

type SchedulePostgresRepository interface {
	GetSchedulers() ([]models.Schedule, error)
	//GetScheduleById(id int) (models.Schedule, error)
	CreateSchedule(schedule models.Schedule) error
	UpdateSchedule(id int, updatedSchedule map[string]interface{}) error
	DeleteSchedule(id int) error
	GetScheduleForGroup(id int) ([]models.Schedule, error)
}

type ScheduleApiRepository interface {
	CreateSchedule(c *gin.Context) error
	GetSchedule(c *gin.Context) error
	GetSchedules(c *gin.Context) error
	UpdateSchedule(c *gin.Context) error
	DeleteSchedule(c *gin.Context) error
	//GetScheduleForGroup(c *gin.Context) error
}
