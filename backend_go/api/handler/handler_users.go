package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler interface {
	CreateStudent(c *gin.Context) error
	GetStudent(c *gin.Context) error
	UpdateStudent(c *gin.Context) error
	DeleteStudent(c *gin.Context) error
}

type userHandler struct {
	logger *zap.Logger
}

func (sh *userHandler) CreateStudent(c *gin.Context) error {
	return nil
}

func (sh *userHandler) GetStudent(c *gin.Context) error {
	return nil
}

func (sh *userHandler) UpdateStudent(c *gin.Context) error {
	return nil
}

func (sh *userHandler) DeleteStudent(c *gin.Context) error {
	return nil
}

func NewStudentHandler() UserHandler {
	return &userHandler{}
}
