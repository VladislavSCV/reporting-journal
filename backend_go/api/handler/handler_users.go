package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/VladislavSCV/internal/database/postgres"
	"github.com/VladislavSCV/internal/model"
	"github.com/VladislavSCV/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler interface {
	Login(c *gin.Context) (model.User, error)
	SignUp(c *gin.Context) error
	GetStudent(c *gin.Context) error
	UpdateStudent(c *gin.Context) error
	DeleteStudent(c *gin.Context) error
}

type userHandler struct {
	logger            *zap.Logger
	servicePostgresql postgres.UserHandlerDB // Сервис для работы с данными студентов

}

func (sh *userHandler) Login(c *gin.Context) (model.User, error) {
	var user model.User
	// принимаем логин и пароль
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return model.User{}, pkg.LogWriteFileReturnError(err)
	}

	// получаем пользователя из базы данных
	userDB, err := sh.servicePostgresql.GetUserByLogin(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return model.User{}, pkg.LogWriteFileReturnError(err)
	}
	// сравниваем пароли
	if userDB.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return model.User{}, pkg.LogWriteFileReturnError(errors.New("invalid credentials"))
	}
	return userDB, nil

}

// SignUp CreateStudent создает нового студента
func (sh *userHandler) SignUp(c *gin.Context) error {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sh.logger.Error("failed to bind student data", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return err
	}

	err := sh.servicePostgresql.CreateUser(&user)
	if err != nil {
		sh.logger.Error("failed to create student", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student"})
		return err
	}

	sh.logger.Info("student created successfully", zap.Int("id", user.ID))
	c.JSON(http.StatusCreated, user)
	return nil
}

// GetStudent получает данные студента по ID
func (sh *userHandler) GetStudent(c *gin.Context) error {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}
	student, err := sh.servicePostgresql.GetUserById(id)
	if err != nil {
		sh.logger.Error("failed to get student", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "student not found"})
		return pkg.LogWriteFileReturnError(err)
	}

	sh.logger.Info("student retrieved successfully", zap.Int("id", student.ID))
	c.JSON(http.StatusOK, student)
	return nil
}

// UpdateStudent обновляет данные студента
func (sh *userHandler) UpdateStudent(c *gin.Context) error {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}
	var updates map[string]string

	if err := c.ShouldBindJSON(&updates); err != nil {
		sh.logger.Error("failed to bind updates", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return err
	}

	err = sh.servicePostgresql.UpdateUser(strconv.Itoa(id), updates)
	if err != nil {
		sh.logger.Error("failed to update student", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update student"})
		return err
	}

	sh.logger.Info("student updated successfully", zap.String("id", strID))
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
	return nil
}

// DeleteStudent удаляет студента по ID
func (sh *userHandler) DeleteStudent(c *gin.Context) error {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return err
	}
	err = sh.servicePostgresql.DeleteUser(id)
	if err != nil {
		sh.logger.Error("failed to delete student", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete student"})
		return err
	}

	sh.logger.Info("student deleted successfully", zap.String("id", strID))
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	return nil
}

// NewStudentHandler создает новый обработчик студентов
func NewStudentHandler(logger *zap.Logger, servicePostgresql postgres.UserHandlerDB) UserHandler {
	return &userHandler{
		logger:            logger,
		servicePostgresql: servicePostgresql,
	}
}
