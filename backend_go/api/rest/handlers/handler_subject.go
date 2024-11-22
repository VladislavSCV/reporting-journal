package handlers

import (
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/subjects"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SubjectHandler struct {
	servicePostgres subjects.SubjectPostgresRepository
}

func (s *SubjectHandler) GetSubjects(c *gin.Context) error {
	subjects, err := s.servicePostgres.GetSubjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch subjects"})
		return err
	}
	c.JSON(http.StatusOK, subjects)
	return nil
}

func (s *SubjectHandler) GetSubjectById(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}
	subject, err := s.servicePostgres.GetSubjectById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch subject"})
		return err
	}

	c.JSON(http.StatusOK, subject)
	return nil
}

func (s *SubjectHandler) CreateSubject(c *gin.Context) error {
	var subject models.Subject

	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return err
	}
	if err := s.servicePostgres.CreateSubject(subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create subject"})
		return err
	}
	c.JSON(http.StatusOK, gin.H{"message": "subject created successfully"})
	return nil
}

func (s *SubjectHandler) UpdateSubject(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}
	var subject map[string]interface{}
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return err
	}
	if err := s.servicePostgres.UpdateSubject(id, subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update subject"})
		return err
	}
	c.JSON(http.StatusOK, gin.H{"message": "subject updated successfully"})
	return nil
}

func (s *SubjectHandler) DeleteSubject(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}
	if err := s.servicePostgres.DeleteSubject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete subject"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "subject deleted successfully"})
	return nil
}

func NewSubjectHandler(servicePostgres subjects.SubjectPostgresRepository) *SubjectHandler {
	return &SubjectHandler{servicePostgres: servicePostgres}
}
