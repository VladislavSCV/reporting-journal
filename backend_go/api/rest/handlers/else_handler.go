package handlers

import (
	"github.com/VladislavSCV/internal/groups"
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/role"
	"github.com/VladislavSCV/internal/users"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"sync"
)

type elseHandler struct {
	usersPostgres users.UserPostgresRepository
	groupPostgres groups.GroupPostgresRepository
	rolePostgres  role.RolePostgresRepository
}

type AdminPanelData struct {
	Users  []models.User  `json:"users"`
	Roles  []models.Role  `json:"roles"`
	Groups []models.Group `json:"groups"`
}

func (h *elseHandler) GetAdminPanelData(c *gin.Context) error {
	var wg sync.WaitGroup
	var users []models.User
	var roles []models.Role
	var groups []models.Group
	errors := make(chan error, 3)

	wg.Add(3)

	// Запускаем горутины для получения данных
	go func() {
		defer wg.Done()
		var err error
		users, err = h.usersPostgres.GetUsers()
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		roles, err = h.rolePostgres.GetRoles()
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		groups, err = h.groupPostgres.GetGroups()
		if err != nil {
			errors <- err
		}
	}()

	wg.Wait()
	close(errors)

	// Если произошла ошибка, возвращаем её клиенту
	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get data"})
		return nil
	}

	// Формируем данные для админ панели
	adminPanelData := AdminPanelData{
		Users:  users,
		Roles:  roles,
		Groups: groups,
	}

	// Логируем данные
	log.Println(adminPanelData)

	// Сериализуем данные в JSON
	data, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(adminPanelData)
	if err != nil {
		// handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize data"})
		return err
	}

	// Отправляем сжатые данные
	c.JSON(http.StatusOK, data)
	return nil
}

func NewElseHandler(usersPostgres users.UserPostgresRepository, groupPostgres groups.GroupPostgresRepository, rolePostgres role.RolePostgresRepository) models.Else {
	return &elseHandler{
		usersPostgres: usersPostgres,
		groupPostgres: groupPostgres,
		rolePostgres:  rolePostgres,
	}
}
