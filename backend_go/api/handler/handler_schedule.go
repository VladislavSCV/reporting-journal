package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ScheduleHandler interface {
	CreateSchedule(c echo.Context) error
	GetSchedule(c echo.Context) error
	UpdateSchedule(c echo.Context) error
	DeleteSchedule(c echo.Context) error
}

type scheduleHandler struct {
	logger *zap.Logger
}

func (sh *scheduleHandler) CreateSchedule(c echo.Context) error {
	return nil
}

func (sh *scheduleHandler) GetSchedule(c echo.Context) error {
	return nil
}

func (sh *scheduleHandler) UpdateSchedule(c echo.Context) error {
	return nil
}

func (sh *scheduleHandler) DeleteSchedule(c echo.Context) error {
	return nil
}

func NewScheduleHandler() ScheduleHandler {
	return &scheduleHandler{}
}
