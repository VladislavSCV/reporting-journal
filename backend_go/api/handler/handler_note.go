package handler

import (
	"go.uber.org/zap"
)

type NoteHandler interface {
	CreateNote(c *gin.Context) error
	GetNote(c *gin.Context) error
	UpdateNote(c *gin.Context) error
	DeleteNote(c *gin.Context) error
}

type noteHandler struct {
	logger *zap.Logger
}

func (noteHandler) CreateNote(c *gin.Context) error {
	return nil
}

func (noteHandler) GetNote(c *gin.Context) error {
	return nil
}

func (noteHandler) UpdateNote(c *gin.Context) error {
	return nil
}

func (noteHandler) DeleteNote(c *gin.Context) error {
	return nil
}

func NewNoteHandler() NoteHandler {
	return &noteHandler{}
}
