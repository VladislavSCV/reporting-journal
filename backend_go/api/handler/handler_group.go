package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type GroupHandler interface {
	CreateGroup(c echo.Context) error
	GetGroup(c echo.Context) error
	UpdateGroup(c echo.Context) error
	DeleteGroup(c echo.Context) error
}

type groupHandler struct {
	logger *zap.Logger
}

func (gh *groupHandler) CreateGroup(c echo.Context) error {
	return nil
}

func (gh *groupHandler) GetGroup(c echo.Context) error {
	return nil
}

func (gh *groupHandler) UpdateGroup(c echo.Context) error {
	return nil
}

func (gh *groupHandler) DeleteGroup(c echo.Context) error {
	return nil
}

func NewGroupHandler() GroupHandler {
	return &groupHandler{}
}
