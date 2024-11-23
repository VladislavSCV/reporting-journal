package handlers

import (
	"net/http"
	"strconv"

	"github.com/VladislavSCV/internal/groups"
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
)

type groupHandler struct {
	dbAndTx groups.GroupPostgresRepository
}

func (gh *groupHandler) GetGroupByID(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return err
	}

	group, err := gh.dbAndTx.GetGroupByID(id)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, group)
	return nil
}

// Создаёт новую группу
func (gh *groupHandler) CreateGroup(c *gin.Context) error {
	var group models.Group

	// Попытка привязать данные из JSON к структуре Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return nil
	}

	// Создание группы через репозиторий
	if err := gh.dbAndTx.CreateGroup(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return err
	}

	c.JSON(http.StatusCreated, group)
	return nil
}

func (gh *groupHandler) GetGroups(c *gin.Context) error {
	groupList, err := gh.dbAndTx.GetAllGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve groups"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"groups": groupList})
	return nil
}

// Обновляет информацию о группе
func (gh *groupHandler) UpdateGroup(c *gin.Context) error {
	var group models.Group

	// Привязка JSON-данных
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return err
	}

	// Обновление группы
	if err := gh.dbAndTx.UpdateGroup(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update group"})
		return err
	}

	c.JSON(http.StatusOK, group)
	return nil
}

// Удаляет группу по ID
func (gh *groupHandler) DeleteGroup(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return err
	}

	// Удаление группы по ID
	if err := gh.dbAndTx.DeleteGroup(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete group"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
	return nil
}

// NewGroupHandler Функция для создания нового GroupHandler
func NewGroupHandler(db groups.GroupPostgresRepository) *groupHandler {
	return &groupHandler{dbAndTx: db}
}
