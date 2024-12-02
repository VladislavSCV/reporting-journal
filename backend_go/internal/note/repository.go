package note

import (
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
)

type NotePostgresRepository interface {
	CreateNote(id int, note models.Note) error
	GetNote(id int) ([]models.Note, error)
	GetGroupNote(id int) ([]models.Note, error)
	GetCuratorGroupNote(teacherId int) ([]models.Note, error)
	GetNotes() (notes []models.Note, err error)
	UpdateNote(id int, newNote map[string]interface{}) error
	DeleteNote(id int) error
}

type NoteApiRepository interface {
	CreateNote(c *gin.Context)
	GetNote(c *gin.Context)
	GetGroupNote(c *gin.Context)
	GetCuratorGroupNote(c *gin.Context)
	GetNotes(c *gin.Context)
	UpdateNote(c *gin.Context)
	DeleteNote(c *gin.Context)
}
