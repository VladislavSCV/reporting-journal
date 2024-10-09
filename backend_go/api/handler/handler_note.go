package handler

import (
	"github.com/gin-gonic/gin"
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

func (nh noteHandler) CreateNote(c *gin.Context) error {
	return nil
}

func (nh noteHandler) GetNote(c *gin.Context) error {
	return nil
}

func (nh noteHandler) UpdateNote(c *gin.Context) error {
	return nil
}

func (nh noteHandler) DeleteNote(c *gin.Context) error {
	return nil
}

func NewNoteHandler() NoteHandler {
	return &noteHandler{}
}
