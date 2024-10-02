package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AuthHandler interface {
	Login(c echo.Context) error
	SignUp(c echo.Context) error
}

type authHandler struct {
	logger zap.Logger
}

func (a *authHandler) Login(c echo.Context) error {
	//TODO Обращаемся к бд и получаем данные пользователя.
	return nil
}

func (a *authHandler) SignUp(c echo.Context) error {
	//TODO Обращаемся к бд и сравниваем данные пользователя
	return nil
}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}
