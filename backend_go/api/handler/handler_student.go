package handler

import (
	"go.uber.org/zap"
)

type StudentHandler interface {
	CreateStudent(c *gin.Context) error
	GetStudent(c *gin.Context) error
	UpdateStudent(c *gin.Context) error
	DeleteStudent(c *gin.Context) error
}

type studentHandler struct {
	logger *zap.Logger
}

func (sh *studentHandler) CreateStudent(c *gin.Context) error {
	return nil
}

func (sh *studentHandler) GetStudent(c *gin.Context) error {
	return nil
}

func (sh *studentHandler) UpdateStudent(c *gin.Context) error {
	return nil
}

func (sh *studentHandler) DeleteStudent(c *gin.Context) error {
	return nil
}

func NewStudentHandler() StudentHandler {
	return &studentHandler{}
}
