package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type StudentHandler interface {
	CreateStudent(c echo.Context) error
	GetStudent(c echo.Context) error
	UpdateStudent(c echo.Context) error
	DeleteStudent(c echo.Context) error
}

type studentHandler struct {
	logger *zap.Logger
}

func (sh *studentHandler) CreateStudent(c echo.Context) error {
	return nil
}

func (sh *studentHandler) GetStudent(c echo.Context) error {
	return nil
}

func (sh *studentHandler) UpdateStudent(c echo.Context) error {
	return nil
}

func (sh *studentHandler) DeleteStudent(c echo.Context) error {
	return nil
}

func NewStudentHandler() StudentHandler {
	return &studentHandler{}
}
