package handlers

import (
	"github.com/VladislavSCV/internal/users"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type roleHandler struct {
	logger            *zap.Logger
	servicePostgresql users.UserPostgresRepository // Сервис для работы с данными пользователей
	serviceRedis      users.UserRedisRepository
}

func (rh *roleHandler) CreateRole(c *gin.Context) error {
	return nil
}

func (rh *roleHandler) GetRole(c *gin.Context) error {
	return nil
}

func (rh *roleHandler) UpdateRole(c *gin.Context) error {
	return nil
}

func (rh *roleHandler) DeleteRole(c *gin.Context) error {
	return nil
}
