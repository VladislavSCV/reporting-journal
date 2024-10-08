package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler interface {
	Login(c *gin.Context) error
	SignUp(c *gin.Context) error
}

type authHandler struct {
	logger zap.Logger
}

func (a *authHandler) Login(c *gin.Context) error {
	//TODO Обращаемся к бд и получаем данные пользователя.
	return nil
}

func (a *authHandler) SignUp(c *gin.Context) error {
	//TODO Обращаемся к бд и сравниваем данные пользователя
	return nil
}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}
