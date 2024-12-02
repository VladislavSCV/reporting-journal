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

func (s *SubjectHandler) GetSubjects(c *gin.Context) {
	subjects, err := s.servicePostgres.GetSubjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch subjects"})
		return
	}
	c.JSON(http.StatusOK, subjects)
	return
}

func (s *SubjectHandler) GetSubjectById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	subject, err := s.servicePostgres.GetSubjectById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch subject"})
		return
	}

	c.JSON(http.StatusOK, subject)
	return
}

func (s *SubjectHandler) CreateSubject(c *gin.Context) {
	var subject models.Subject

	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := s.servicePostgres.CreateSubject(subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create subject"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "subject created successfully"})
	return
}

func (s *SubjectHandler) UpdateSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var subject map[string]interface{}
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := s.servicePostgres.UpdateSubject(id, subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update subject"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "subject updated successfully"})
	return
}

func (s *SubjectHandler) DeleteSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := s.servicePostgres.DeleteSubject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete subject"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "subject deleted successfully"})
	return
}

func NewSubjectHandler(servicePostgres subjects.SubjectPostgresRepository) *SubjectHandler {
	return &SubjectHandler{servicePostgres: servicePostgres}
}
