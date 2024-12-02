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
	CreateSubject(c *gin.Context)
	GetSubjects(c *gin.Context)
	GetSubjectById(c *gin.Context)
	UpdateSubject(c *gin.Context)
	DeleteSubject(c *gin.Context)
}
