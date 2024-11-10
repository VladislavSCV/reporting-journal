package handlers

import (
	"errors"
	"fmt"
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
		sh.logger.Error("failed to bind user data", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return pkg.LogWriteFileReturnError(errors.New("invalid input"))
	}

	// Получаем пользователя из базы данных
	userDB, err := sh.servicePostgresql.GetUserByLogin(user.Login)
	if err != nil {
		sh.logger.Error("failed to get user from PostgreSQL",
			zap.String("login", user.Login),
			zap.Error(err),
		)
		c.Status(http.StatusNotFound)
		return pkg.LogWriteFileReturnError(errors.New("failed to get user"))
	}

	// Проверка пароля
	isValid, err := pkg.VerifyPassword(user.Hash, userDB.Salt, userDB.Hash)
	if err != nil {
		sh.logger.Error("error verifying password",
			zap.String("login", user.Login),
			zap.Error(err),
		)
		c.Status(http.StatusInternalServerError)
		return pkg.LogWriteFileReturnError(errors.New("error verifying password"))
	}

	if !isValid {
		sh.logger.Info("invalid credentials",
			zap.String("login", user.Login),
		)
		c.Status(http.StatusUnauthorized)
		return pkg.LogWriteFileReturnError(errors.New("invalid credentials"))
	}

	// Генерация токена
	token, err := pkg.GenerateJWT(userDB.ID)
	if err != nil {
		sh.logger.Error("failed to generate token",
			zap.Int("id", userDB.ID),
			zap.Error(err),
		)
		c.Status(http.StatusInternalServerError)
		return pkg.LogWriteFileReturnError(errors.New("failed to generate token"))
	}

	// Успешная аутентификация
	c.JSON(http.StatusOK, gin.H{"user": userDB, "token": token})
	return nil
}

// SignUp (Регистрация) создает нового студента
func (sh *userHandler) SignUp(c *gin.Context) error {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sh.logger.Error("invalid input",
			zap.Error(err),
		)
		fmt.Println(user)
		c.Status(http.StatusBadRequest)
		return pkg.LogWriteFileReturnError(errors.New("invalid input"))
	}

	fmt.Println(user)

	if err := sh.servicePostgresql.CreateUser(&user); err != nil {
		sh.logger.Error("failed to create user",
			zap.Int("id", user.ID),
			zap.String("login", user.Login),
			zap.Error(err),
		)
		c.Status(http.StatusInternalServerError)
		return pkg.LogWriteFileReturnError(errors.New("failed to create user"))
	}

	token, err := pkg.GenerateJWT(user.ID)
	if err != nil {
		sh.logger.Error("failed to generate token",
			zap.Int("id", user.ID),
			zap.Error(err),
		)
		c.Status(http.StatusInternalServerError)
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

	user, err := sh.servicePostgresql.GetUserById(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return pkg.LogWriteFileReturnError(errors.New("failed to retrieve user from PostgreSQL"))
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

func (sh *userHandler) VerifyToken(c *gin.Context) error {
	var request struct {
		Token string `json:"token"` // Убедитесь, что тело запроса имеет нужную структуру
	}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return err
	}

	_, err = pkg.VerifyToken(request.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return pkg.LogWriteFileReturnError(errors.New("invalid token"))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
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
