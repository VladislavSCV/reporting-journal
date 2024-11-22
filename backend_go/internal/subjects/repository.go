package subjects

import (
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
)

type SubjectPostgresRepository interface {
	CreateSubject(subject models.Subject) error
	GetSubjects() ([]models.Subject, error)
	GetSubjectById(id int) (models.Subject, error)
	UpdateSubject(id int, UpdatedSubjects map[string]interface{}) error
	DeleteSubject(id int) error
}

type SubjectApiRepository interface {
	CreateSubject(c *gin.Context) error
	GetSubjects(c *gin.Context) error
	GetSubjectById(c *gin.Context) error
	UpdateSubject(c *gin.Context) error
	DeleteSubject(c *gin.Context) error
}
