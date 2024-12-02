package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/note"
	"github.com/gin-gonic/gin"
)

type noteHandler struct {
	servicePostgres note.NotePostgresRepository
}

// CreateNote добавляет новую заметку
func (nh *noteHandler) CreateNote(c *gin.Context) error {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}
	var modelsNote models.Note
	if err := c.ShouldBindJSON(&modelsNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return err
	}

	if err := nh.servicePostgres.CreateNote(id, modelsNote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create note"})
		return err
	}

	c.JSON(http.StatusCreated, gin.H{"message": "note created successfully"})
	return nil
}

// GetNote получает заметку по ID
func (nh *noteHandler) GetNote(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}

	note, err := nh.servicePostgres.GetNote(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch note"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"notes": note})
	return nil
}

func (nh *noteHandler) GetGroupNote(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}

	note, err := nh.servicePostgres.GetGroupNote(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch note"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"notes": note})
	return nil
}

func (nh *noteHandler) GetCuratorGroupNote(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}

	note, err := nh.servicePostgres.GetCuratorGroupNote(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch note"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"notes": note})
	return nil
}

// GetNotes получает список всех заметок
func (nh *noteHandler) GetNotes(c *gin.Context) error {
	notes, err := nh.servicePostgres.GetNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch notes"})
		return err
	}

	c.JSON(http.StatusOK, notes)
	return nil
}

// UpdateNote обновляет данные заметки
func (nh *noteHandler) UpdateNote(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}

	var newNote map[string]interface{}
	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return err
	}
	log.Println(newNote)
	if err := nh.servicePostgres.UpdateNote(id, newNote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update note"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "note updated successfully"})
	return nil
}

// DeleteNote удаляет заметку по ID
func (nh *noteHandler) DeleteNote(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}

	if err := nh.servicePostgres.DeleteNote(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete note"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "note deleted successfully"})
	return nil
}

// NewNoteHandler создает экземпляр noteHandler
func NewNoteHandler(service note.NotePostgresRepository) *noteHandler {
	return &noteHandler{
		servicePostgres: service,
	}
}
