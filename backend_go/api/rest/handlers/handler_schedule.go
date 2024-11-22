package handlers

import (
	"github.com/VladislavSCV/pkg"
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
func (sh *scheduleHandler) CreateSchedule(c *gin.Context) error {
	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return pkg.LogWriteFileReturnError(err)
	}

	if err := sh.repository.CreateSchedule(schedule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule: " + err.Error()})
		return pkg.LogWriteFileReturnError(err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Schedule created successfully"})
	return nil
}

// GetSchedule получает расписание по ID
func (sh *scheduleHandler) GetSchedule(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return pkg.LogWriteFileReturnError(err)
	}

	schedule, err := sh.repository.GetScheduleById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return pkg.LogWriteFileReturnError(err)
	}

	c.JSON(http.StatusOK, schedule)
	return nil
}

func (sh *scheduleHandler) GetSchedules(c *gin.Context) error {
	schedulers, err := sh.repository.GetSchedulers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Scheduler not found"})
		return err
	}
	c.JSON(http.StatusOK, schedulers)
	return nil
}

// UpdateSchedule обновляет расписание по ID
func (sh *scheduleHandler) UpdateSchedule(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return pkg.LogWriteFileReturnError(err)
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return pkg.LogWriteFileReturnError(err)
	}

	if err := sh.repository.UpdateSchedule(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule: " + err.Error()})
		return pkg.LogWriteFileReturnError(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule updated successfully"})
	return nil
}

// DeleteSchedule удаляет расписание по ID
func (sh *scheduleHandler) DeleteSchedule(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return pkg.LogWriteFileReturnError(err)
	}

	if err := sh.repository.DeleteSchedule(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule: " + err.Error()})
		return pkg.LogWriteFileReturnError(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
	return nil
}
