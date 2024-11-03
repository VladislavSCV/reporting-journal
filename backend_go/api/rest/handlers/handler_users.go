package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/users"
	"github.com/VladislavSCV/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userHandler struct {
	logger            *zap.Logger
	servicePostgresql users.UserPostgresRepository // Сервис для работы с данными пользователей
	serviceRedis      users.UserRedisRepository
}

// Login аутентифицирует пользователя
//
//	Accepts:	JSON {login: string, password: string}
//	Returns:	JSON {user: models.User}
//	Returns error:	invalid input, failed to get user, invalid credentials
func (sh *userHandler) Login(c *gin.Context) error {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return pkg.LogWriteFileReturnError(err)
	}

	userDB, err := sh.servicePostgresql.GetUserByLogin(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return pkg.LogWriteFileReturnError(err)
	}

	if userDB.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return pkg.LogWriteFileReturnError(errors.New("invalid credentials"))
	}

	err = sh.serviceRedis.SaveInCache(&userDB)
	if err != nil {
		return pkg.LogWriteFileReturnError(errors.New("failed to set user in cache"))
	}

	token, err := pkg.GenerateJWT(userDB.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return pkg.LogWriteFileReturnError(err)
	}

	// Успешная аутентификация, возвращаем пользователя и JWT
	c.JSON(http.StatusOK, gin.H{"user": userDB, "token": token})
	return nil
}

// SignUp Регистрация. создает нового студента
func (sh *userHandler) SignUp(c *gin.Context) error {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sh.logger.Error("failed to bind user data", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return err
	}

	err := sh.servicePostgresql.CreateUser(&user)
	if err != nil {
		sh.logger.Error("failed to create user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return err
	}

	err = sh.serviceRedis.SaveInCache(&user)
	if err != nil {
		return err
	}

	token, err := pkg.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return pkg.LogWriteFileReturnError(err)
	}

	sh.logger.Info("user created successfully", zap.Int("id", user.ID))
	c.JSON(http.StatusCreated, gin.H{"user": user, "token": token})
	return nil
}

// GetUsers возвращает список всех студентов
//
//	@return error - ошибка, если она возникла
func (sh *userHandler) GetUsers(c *gin.Context) error {
	usersDB, err := sh.servicePostgresql.GetUsers()
	if err != nil {
		pkg.LogWriteFileReturnError(errors.New("failed to retrieve users from the database"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return err
	}
	c.JSON(http.StatusOK, gin.H{"users": usersDB})
	return err
}

// GetStudent получает данные студента по ID
func (sh *userHandler) GetUser(c *gin.Context) error {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		pkg.LogWriteFileReturnError(errors.New("invalid user ID format"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return err
	}

	user, err := sh.serviceRedis.GetUserById(id)
	if err != nil {
		pkg.LogWriteFileReturnError(errors.New("failed to retrieve user from Redis"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return err
	}
	if user.Name == "" {
		user, err = sh.servicePostgresql.GetUserById(id)
		if err != nil {
			pkg.LogWriteFileReturnError(errors.New("failed to retrieve user from PostgreSQL"))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
			return err
		}
		user.ID = id
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
	return err
}

// GetUserByLogin возвращает данные студента по логину
//
//	Accepts:	JSON {login: string}
//	Returns:	JSON {user: models.User}
//	Returns error:	invalid input, failed to get user
func (sh *userHandler) GetUserByLogin(c *gin.Context) (models.User, error) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sh.logger.Error("failed to bind student data", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return models.User{}, err
	}

	userFromDB, err := sh.servicePostgresql.GetUserByLogin(user.Login)
	if err != nil {
		pkg.LogWriteFileReturnError(errors.New("failed to retrieve user from PostgreSQL"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return models.User{}, err
	}

	err = sh.serviceRedis.SaveInCache(&user)
	if err != nil {
		return models.User{}, err
	}
	return userFromDB, nil
}

// UpdateStudent обновляет данные студента
func (sh *userHandler) UpdateUser(c *gin.Context) error {
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

// DeleteUser удаляет студента по ID
func (sh *userHandler) DeleteUser(c *gin.Context) error {
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
func NewUserHandler(servicePostgresql users.UserPostgresRepository, serviceRedis users.UserRedisRepository) users.UserAPIRepository {
	logger, err := zap.NewProduction()
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}
	return &userHandler{
		logger:            logger,
		servicePostgresql: servicePostgresql,
		serviceRedis:      serviceRedis,
	}
}
