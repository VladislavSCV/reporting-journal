package handlers

import (
	"net/http"
	"strconv"

	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/role"
	"github.com/VladislavSCV/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type roleHandler struct {
	logger            *zap.Logger
	servicePostgresql role.RolePostgresRepository // Сервис для работы с данными пользователей
}

func (rh *roleHandler) CreateRole(c *gin.Context) error {
	var modelRole models.Role
	err := c.ShouldBindJSON(&modelRole)
	if err != nil {
		return err
	}
	err = rh.servicePostgresql.CreateRole(&modelRole)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return pkg.LogWriteFileReturnError(err)
	}

	return nil
}

func (rh *roleHandler) GetRoles(c *gin.Context) error {
	roles, err := rh.servicePostgresql.GetRoles()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return pkg.LogWriteFileReturnError(err)
	}
	c.JSON(http.StatusOK, roles)
	return nil
}

func (rh *roleHandler) GetRole(c *gin.Context) error {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return pkg.LogWriteFileReturnError(err)
	}

	getRole, err := rh.servicePostgresql.GetRole(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return pkg.LogWriteFileReturnError(err)
	}
	c.JSON(http.StatusOK, getRole)
	return nil
}

func (rh *roleHandler) DeleteRole(c *gin.Context) error {
	//var modelRole models.Role
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return pkg.LogWriteFileReturnError(err)
	}
	err = rh.servicePostgresql.DeleteRole(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return pkg.LogWriteFileReturnError(err)
	}

	return nil
}

func NewRoleHandler(servicePostgresql role.RolePostgresRepository) role.RoleApiRepository {
	return &roleHandler{
		servicePostgresql: servicePostgresql,
	}
}
