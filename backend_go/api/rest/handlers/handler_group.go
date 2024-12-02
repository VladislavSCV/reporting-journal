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

func (gh *groupHandler) GetGroupByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	group, err := gh.dbAndTx.GetGroupByID(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, group)
	return
}

// Создаёт новую группу
func (gh *groupHandler) CreateGroup(c *gin.Context) {
	var group models.Group

	// Попытка привязать данные из JSON к структуре Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Создание группы через репозиторий
	id, err := gh.dbAndTx.CreateGroup(&group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return
	}

	group.Id = id

	c.JSON(http.StatusCreated, group)
	return
}

func (gh *groupHandler) GetGroups(c *gin.Context) {
	groupList, err := gh.dbAndTx.GetGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve groups"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"groups": groupList})
	return
}

// Обновляет информацию о группе
func (gh *groupHandler) UpdateGroup(c *gin.Context) {
	var group models.Group

	// Привязка JSON-данных
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Обновление группы
	if err := gh.dbAndTx.UpdateGroup(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update group"})
		return
	}

	c.JSON(http.StatusOK, group)
	return
}

// Удаляет группу по ID
func (gh *groupHandler) DeleteGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	// Удаление группы по ID
	if err := gh.dbAndTx.DeleteGroup(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete group"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
	return
}

// NewGroupHandler Функция для создания нового GroupHandler
func NewGroupHandler(db groups.GroupPostgresRepository) *groupHandler {
	return &groupHandler{dbAndTx: db}
}
