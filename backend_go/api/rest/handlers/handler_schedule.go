package handlers

import (
	"net/http"
	"strconv"

	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/schedules"
	"github.com/gin-gonic/gin"
)

type scheduleHandler struct {
	repository schedules.SchedulePostgresRepository
}

// NewScheduleHandler создает новый экземпляр scheduleHandler
func NewScheduleHandler(repo schedules.SchedulePostgresRepository) *scheduleHandler {
	return &scheduleHandler{repository: repo}
}

// CreateSchedule создает новое расписание
func (sh *scheduleHandler) CreateSchedule(c *gin.Context) {
	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if err := sh.repository.CreateSchedule(schedule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Schedule created successfully"})
	return
}

// GetSchedule получает расписание по ID
func (sh *scheduleHandler) GetSchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	schedule, err := sh.repository.GetScheduleForGroup(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"schedule": schedule})
	return
}

func (sh *scheduleHandler) GetSchedules(c *gin.Context) {
	schedulers, err := sh.repository.GetSchedulers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Scheduler not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"schedulers": schedulers})
	return
}

// UpdateSchedule обновляет расписание по ID
func (sh *scheduleHandler) UpdateSchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if err := sh.repository.UpdateSchedule(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule updated successfully"})
	return
}

// DeleteSchedule удаляет расписание по ID
func (sh *scheduleHandler) DeleteSchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := sh.repository.DeleteSchedule(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
	return
}

//func (sh *scheduleHandler) GetScheduleForGroup(c *gin.Context) error {
//
//	group, err := sh.repository.GetScheduleForGroup()
//	if err != nil {
//		return err
//	}
//}
