package handlers

import (
	"errors"
	"github.com/VladislavSCV/internal/attendance"
	"github.com/VladislavSCV/internal/groups"
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/role"
	"github.com/VladislavSCV/internal/users"
	"github.com/VladislavSCV/pkg"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type elseHandler struct {
	usersPostgres      users.UserPostgresRepository
	groupPostgres      groups.GroupPostgresRepository
	rolePostgres       role.RolePostgresRepository
	attendancePostgres attendance.AttendancePostgresRepository
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

func (h *elseHandler) GetCuratorGroupsStudentList(c *gin.Context) error {
	// Извлечение токена из заголовка Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		//sh.logger.Error("missing Authorization header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Authorization header"})
		return errors.New("missing Authorization header")
	}

	// Проверка формата заголовка, например: "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		//sh.logger.Error("invalid Authorization header format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization header format"})
		return errors.New("invalid Authorization header format")
	}
	token := parts[1]

	id, roleId, err := pkg.ParseJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return err
	}

	if roleId != 2 && roleId != 3 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		return errors.New("access denied")
	}

	//groups, err := h.groupPostgres.GetGroups()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get groups"})
	//	return err
	//}

	groups, err := h.groupPostgres.GetCuratorGroups(id)
	if err != nil {
		return err
	}

	// Сериализуем данные в JSON
	data, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(groups)
	if err != nil {
		// handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize data"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"groups": data})

	return nil
}
func (h *elseHandler) StudentsAttendance(c *gin.Context) error {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return err
	}

	users, err := h.usersPostgres.GetUsersByGroupID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return err
	}

	// Сериализуем данные в JSON
	data, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(users)
	if err != nil {
		// handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize data"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"users": data})
	return nil
}

type UpdateAttendanceRequest struct {
	StudentId int    `json:"studentId"`
	Status    string `json:"status"`
}

func (h *elseHandler) UpdateAttendance(c *gin.Context) error {
	var updateAttendanceRequest UpdateAttendanceRequest
	err := c.ShouldBindJSON(&updateAttendanceRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return err
	}

	err = h.attendancePostgres.UpdateAttendance(updateAttendanceRequest.StudentId, updateAttendanceRequest.Status)
	if err != nil {
		err = h.attendancePostgres.AddAttendance(updateAttendanceRequest.StudentId, updateAttendanceRequest.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attendance"})
			return err
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance updated successfully"})
	return nil

}

func NewElseHandler(usersPostgres users.UserPostgresRepository, groupPostgres groups.GroupPostgresRepository, rolePostgres role.RolePostgresRepository, attendancePostgres attendance.AttendancePostgresRepository) models.Else {
	return &elseHandler{
		usersPostgres:      usersPostgres,
		groupPostgres:      groupPostgres,
		rolePostgres:       rolePostgres,
		attendancePostgres: attendancePostgres,
	}
}
