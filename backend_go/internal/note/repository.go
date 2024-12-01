package note

import (
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
)

type NotePostgresRepository interface {
	CreateNote(note models.Note) error
	GetNote(id int) ([]models.Note, error)
	GetGroupNote(id int) ([]models.Note, error)
	GetCuratorGroupNote(teacherId int) ([]models.Note, error)
	GetNotes() (notes []models.Note, err error)
	UpdateNote(id int, newNote map[string]interface{}) error
	DeleteNote(id int) error
}

type NoteApiRepository interface {
	CreateNote(c *gin.Context) error
	GetNote(c *gin.Context) error
	GetGroupNote(c *gin.Context) error
	GetCuratorGroupNote(c *gin.Context) error
	GetNotes(c *gin.Context) error
	UpdateNote(c *gin.Context) error
	DeleteNote(c *gin.Context) error
}
