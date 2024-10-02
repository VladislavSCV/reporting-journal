package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type NoteHandler interface {
	CreateNote(c echo.Context) error
	GetNote(c echo.Context) error
	UpdateNote(c echo.Context) error
	DeleteNote(c echo.Context) error
}

type noteHandler struct {
	logger *zap.Logger
}

func (noteHandler) CreateNote(c echo.Context) error {
	return nil
}

func (noteHandler) GetNote(c echo.Context) error {
	return nil
}

func (noteHandler) UpdateNote(c echo.Context) error {
	return nil
}

func (noteHandler) DeleteNote(c echo.Context) error {
	return nil
}

func NewNoteHandler() NoteHandler {
	return &noteHandler{}
}
