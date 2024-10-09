package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GroupHandler interface {
	CreateGroup(c *gin.Context) error
	GetGroup(c *gin.Context) error
	UpdateGroup(c *gin.Context) error
	DeleteGroup(c *gin.Context) error
}

type groupHandler struct {
	logger *zap.Logger
}

func (gh groupHandler) CreateGroup(c *gin.Context) error {
	return nil
}

func (gh groupHandler) GetGroup(c *gin.Context) error {
	return nil
}

func (gh groupHandler) UpdateGroup(c *gin.Context) error {
	return nil
}

func (gh groupHandler) DeleteGroup(c *gin.Context) error {
	return nil
}

func NewGroupHandler() GroupHandler {
	return &groupHandler{}
}
