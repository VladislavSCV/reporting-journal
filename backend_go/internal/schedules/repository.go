package schedules

import (
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
)

type SchedulePostgresRepository interface {
	GetSchedulers() ([]models.Schedule, error)
	//GetScheduleById(id int) (models.Schedule, error)
	CreateSchedule(schedule models.CreateSchedule) error
	UpdateSchedule(id int, updatedSchedule map[string]interface{}) error
	DeleteSchedule(id int) error
	GetScheduleForGroup(id int) ([]models.Schedule, error)
}

type ScheduleApiRepository interface {
	CreateSchedule(c *gin.Context)
	GetSchedule(c *gin.Context)
	GetSchedules(c *gin.Context)
	UpdateSchedule(c *gin.Context)
	DeleteSchedule(c *gin.Context)
	//GetScheduleForGroup(c *gin.Context) error
}
