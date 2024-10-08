package handler

import (
	"go.uber.org/zap"
)

type ScheduleHandler interface {
	CreateSchedule(c *gin.Context) error
	GetSchedule(c *gin.Context) error
	UpdateSchedule(c *gin.Context) error
	DeleteSchedule(c *gin.Context) error
}

type scheduleHandler struct {
	logger *zap.Logger
}

func (sh *scheduleHandler) CreateSchedule(c *gin.Context) error {
	return nil
}

func (sh *scheduleHandler) GetSchedule(c *gin.Context) error {
	return nil
}

func (sh *scheduleHandler) UpdateSchedule(c *gin.Context) error {
	return nil
}

func (sh *scheduleHandler) DeleteSchedule(c *gin.Context) error {
	return nil
}

func NewScheduleHandler() ScheduleHandler {
	return &scheduleHandler{}
}
