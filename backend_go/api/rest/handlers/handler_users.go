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
	if err := c.ShouldBindJSON(&user); err != nil {
		return pkg.LogWriteFileReturnError(errors.New("invalid input"))
	}

	userDB, err := sh.servicePostgresql.GetUserByLogin(user.Login)
	if err != nil {
		return pkg.LogWriteFileReturnError(errors.New("failed to get user"))
	}

	if userDB.Password != user.Password {
		return pkg.LogWriteFileReturnError(errors.New("invalid credentials"))
	}

	token, err := pkg.GenerateJWT(userDB.ID)
	if err != nil {
		return pkg.LogWriteFileReturnError(errors.New("failed to generate token"))
	}

	// Успешная аутентификация, возвращаем пользователя и JWT
	c.JSON(http.StatusOK, gin.H{"user": userDB, "token": token})
	return nil
}

// SignUp (Регистрация) создает нового студента
func (sh *userHandler) SignUp(c *gin.Context) error {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return pkg.LogWriteFileReturnError(errors.New("invalid input"))
	}

	if err := sh.servicePostgresql.CreateUser(&user); err != nil {
		return pkg.LogWriteFileReturnError(errors.New("failed to create user"))
	}

	token, err := pkg.GenerateJWT(user.ID)
	if err != nil {
		return pkg.LogWriteFileReturnError(errors.New("failed to generate token"))
	}

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
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return pkg.LogWriteFileReturnError(errors.New("invalid user ID format"))
	}

	user, err := sh.serviceRedis.GetUserById(id)
	if err != nil || (user.Name == "") {
		user, err = sh.servicePostgresql.GetUserById(id)
		if err != nil {
			return pkg.LogWriteFileReturnError(errors.New("failed to retrieve user from PostgreSQL"))
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
	return nil
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
		return pkg.LogWriteFileReturnError(errors.New("invalid id"))
	}
	var updates map[string]string

	if err := c.ShouldBindJSON(&updates); err != nil {
		return pkg.LogWriteFileReturnError(errors.New("invalid input"))
	}

	if err := sh.servicePostgresql.UpdateUser(strconv.Itoa(id), updates); err != nil {
		return pkg.LogWriteFileReturnError(errors.New("failed to update user"))
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
	return nil
}

// DeleteUser удаляет студента по ID
func (sh *userHandler) DeleteUser(c *gin.Context) error {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return pkg.LogWriteFileReturnError(errors.New("invalid id"))
	}
	if err := sh.servicePostgresql.DeleteUser(id); err != nil {
		return pkg.LogWriteFileReturnError(errors.New("failed to delete user"))
	}

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
