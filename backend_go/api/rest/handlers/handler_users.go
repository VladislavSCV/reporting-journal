package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"fmt"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/users"
	"github.com/VladislavSCV/pkg"
)

type userHandler struct {
	logger            *zap.Logger
	servicePostgresql users.UserPostgresRepository // Сервис для работы с данными пользователей
	// serviceRedis      users.UserRedisRepository
}

// Login аутентифицирует пользователя
//
//	Accepts:	JSON {login: string, password: string}
//	Returns:	JSON {user: models.User}
//	Returns error:	invalid input, failed to get user, invalid credentials
func (sh *userHandler) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sh.logger.Error("failed to bind user data", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return
	}

	// Получаем пользователя из базы данных
	userDB, err := sh.servicePostgresql.GetUserByLogin(user.Login)
	if err != nil {
		sh.logger.Error("failed to get user from PostgreSQL",
			zap.String("login", user.Login),
			zap.Error(err),
		)
		c.Status(http.StatusUnauthorized)
		return
	}

	// TODO раскомментировать код ниже
	// Проверка пароля
	//isValid, err := pkg.VerifyPassword(user.Hash, userDB.Salt, userDB.Hash)
	//if err != nil {
	//	sh.logger.Error("error verifying password",
	//		zap.String("login", user.Login),
	//		zap.String("password", user.Salt),
	//		zap.String("hash", userDB.Hash),
	//		zap.Error(err),
	//	)
	//	c.Status(http.StatusInternalServerError)
	//	return
	//}
	//
	//if !isValid {
	//	sh.logger.Info("invalid credentials",
	//		zap.String("login", user.Login),
	//	)
	//	c.Status(http.StatusUnauthorized)
	//	return
	//}

	// Генерация токена
	token, err := pkg.GenerateJWT(userDB.ID, userDB.RoleID)
	if err != nil {
		sh.logger.Error("failed to generate token",
			zap.Int("id", userDB.ID),
			zap.Error(err),
		)
		c.Status(http.StatusInternalServerError)
		return
	}

	_ = token
	//sh.servicePostgresql.UpdateToken(userDB.ID, token)

	//log.Println(userDB)
	// Успешная аутентификация
	fmt.Println("-----------------------------------------------------------", userDB)
	c.JSON(http.StatusOK, gin.H{"user": userDB, "token": token})
}

//func (sh *userHandler) GetUserRole(token string) (string, error) {
//	user, err := sh.servicePostgresql.GetUserByToken(token)
//	if err != nil {
//		return "", err
//	}
//
//	return user, nil
//}

// SignUp (Регистрация) создает нового студента
func (sh *userHandler) SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sh.logger.Error("invalid input",
			zap.Error(err),
		)
		c.Status(http.StatusBadRequest)
		return
	}

	if user.RoleID == 1 {
		responseUser, token, err := sh.servicePostgresql.CreateStudent(&user)
		if errors.Is(err, errors.New("пользователь с таким логином уже существует")) {
			c.Status(http.StatusConflict)
			return
		}
		if err != nil {
			sh.logger.Error("failed to create user",
				zap.Int("id", responseUser.ID),
				zap.String("login", responseUser.Login),
				zap.Error(err),
			)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"user": responseUser, "token": token})
		return
	} else {
		responseUser, token, err := sh.servicePostgresql.CreateTeacher(&user)
		if errors.Is(err, errors.New("пользователь с таким логином уже существует")) {
			c.Status(http.StatusConflict)
			return
		}
		if err != nil {
			sh.logger.Error("failed to create user",
				zap.Int("id", responseUser.ID),
				zap.String("login", responseUser.Login),
				zap.Error(err),
			)
			c.Status(http.StatusInternalServerError)
			return
		}

		log.Println()

		c.JSON(http.StatusCreated, gin.H{"user": responseUser, "token": token})
		return
	}
}

// GetUsers возвращает список всех студентов
//
//	@return error - ошибка, если она возникла
func (sh *userHandler) GetUsers(c *gin.Context) {
	usersDB, err := sh.servicePostgresql.GetUsers()
	if err != nil {
		pkg.LogWriteFileReturnError(errors.New("failed to retrieve users from the database"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": usersDB})
	return
}

// GetStudent получает данные студента по ID
func (sh *userHandler) GetUser(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return
	}

	user, err := sh.servicePostgresql.GetUserById(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
	return
}

func (sh *userHandler) GetStudents(c *gin.Context) {
	students, err := sh.servicePostgresql.GetStudents()
	if err != nil {
		pkg.LogWriteFileReturnError(errors.New("failed to retrieve users from the database"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"students": students})
	return
}

func (sh *userHandler) GetTeachers(c *gin.Context) {
	teachers, err := sh.servicePostgresql.GetTeachers()
	if err != nil {
		pkg.LogWriteFileReturnError(errors.New("failed to retrieve users from the database"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"teachers": teachers})
	return
}

func (sh *userHandler) GetUserByToken(c *gin.Context) {
	// Извлечение токена из заголовка Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		//sh.logger.Error("missing Authorization header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Authorization header"})
		return
	}

	// Проверка формата заголовка, например: "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		//sh.logger.Error("invalid Authorization header format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization header format"})
		return
	}
	token := parts[1]

	id, _, err := pkg.ParseJWT(token)
	if err != nil {
		return
	}

	userFromDB, err := sh.servicePostgresql.GetUserById(id)
	if err != nil {
		sh.logger.Error("failed to retrieve user from PostgreSQL",
			zap.String("token", token),
			zap.Error(err),
		)
		pkg.LogWriteFileReturnError(errors.New("failed to retrieve user from PostgreSQL"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userFromDB})
	sh.logger.Info("successfully retrieved user from the database",
		zap.String("token", token),
	)

	return
}

// GetUserByLogin возвращает данные студента по логину
//
//	Accepts:	JSON {login: string}
//	Returns:	JSON {user: models.User}
//	Returns error:	invalid input, failed to get user
func (sh *userHandler) GetUserByLogin(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sh.logger.Error("failed to bind student data", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userFromDB, err := sh.servicePostgresql.GetUserByLogin(user.Login)
	if err != nil {
		pkg.LogWriteFileReturnError(errors.New("failed to retrieve user from PostgreSQL"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	// err = sh.serviceRedis.SaveInCache(&user)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userFromDB})
	return
}

// UpdateStudent обновляет данные студента
func (sh *userHandler) UpdateUser(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return
	}
	var updates map[string]string

	if err := c.ShouldBindJSON(&updates); err != nil {
		return
	}

	if err := sh.servicePostgresql.UpdateUser(id, updates); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
	return
}

// DeleteUser удаляет студента по ID
func (sh *userHandler) DeleteUser(c *gin.Context) {
	if sh.logger == nil {
		return
	}

	strID := c.Param("id")
	id, err := strconv.Atoi(strID)

	if err != nil {
		sh.logger.Error("invalid user ID format", zap.Error(err))
		return
	}

	if sh.servicePostgresql == nil {
		sh.logger.Error("nil servicePostgresql")
		return
	}

	sh.logger.Info("deleting user", zap.Int("user_id", id))

	if err := sh.servicePostgresql.DeleteUser(id); err != nil {
		sh.logger.Error("failed to delete user", zap.Int("user_id", id), zap.Error(err))
		return
	}

	sh.logger.Info("successfully deleted user", zap.Int("user_id", id))
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	return
}

func (sh *userHandler) VerifyToken(c *gin.Context) {
	// Извлечение токена из заголовка Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		sh.logger.Error("missing Authorization header")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Authorization header"})
		return
	}

	// Проверка формата заголовка, например: "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		sh.logger.Error("invalid Authorization header format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization header format"})
		return
	}
	token := parts[1]

	// Парсинг JWT токена
	userID, userRoleId, err := pkg.ParseJWT(token)
	if err != nil {
		sh.logger.Error("failed to parse JWT",
			zap.Error(err),
		)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if userRoleId == 0 {
		sh.logger.Error("empty user ID from JWT")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Лог успешного выполнения
	sh.logger.Info("successfully parsed JWT",
		zap.Int("user_role_id", userRoleId),
	)

	// Ответ пользователю
	c.JSON(http.StatusOK, gin.H{
		"id":      userID,
		"role_id": userRoleId,
	})
	return
}

// NewStudentHandler создает новый обработчик студентов
func NewUserHandler(servicePostgresql users.UserPostgresRepository) users.UserAPIRepository {
	logger, err := zap.NewProduction()
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}
	return &userHandler{
		logger:            logger,
		servicePostgresql: servicePostgresql,
		// serviceRedis:      serviceRedis,
	}
}
