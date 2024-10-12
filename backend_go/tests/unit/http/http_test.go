package http

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/VladislavSCV/api/handler"
	"github.com/VladislavSCV/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserHandlerDB - это мока для UserHandlerDB
type MockUserHandlerDB struct {
	mock.Mock
}

func (m *MockUserHandlerDB) GetUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserHandlerDB) GetUserByLogin(login string) (model.User, error) {
	args := m.Called(login)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserHandlerDB) GetUserById(id int) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserHandlerDB) CreateUser(user model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserHandlerDB) UpdateUser(id int, updates map[string]string) error {
	args := m.Called(id, updates)
	return args.Error(0)
}

func (m *MockUserHandlerDB) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestLogin(t *testing.T) {
	// Настройка
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockDB := new(MockUserHandlerDB)
	//logger := zap.NewNop()

	authHandler := handler.NewAuthHandler()

	r.POST("/login", func(c *gin.Context) {
		authHandler.Login(c)
	})

	// Пример успешного входа
	mockDB.On("GetUserByLogin", "testuser").Return(model.User{Login: "testuser"}, nil)

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"login":"testuser", "password":"testpass"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)

	// Пример неверного входа
	mockDB.On("GetUserByLogin", "wronguser").Return(model.User{}, sql.ErrNoRows)

	req, _ = http.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"login":"wronguser", "password":"wrongpass"}`))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockDB.AssertExpectations(t)
}

func TestSignUp(t *testing.T) {
	// Настройка
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockDB := new(MockUserHandlerDB)
	//logger := zap.NewNop()

	authHandler := handler.NewAuthHandler()
	r.POST("/signup", func(c *gin.Context) {
		authHandler.SignUp(c)
	})

	// Пример успешной регистрации
	newUser := model.User{Name: "Test User", Login: "testuser", Password: "testpass"}
	mockDB.On("CreateUser", newUser).Return(nil)

	req, _ := http.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"Test User", "login":"testuser", "password":"testpass"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockDB.AssertExpectations(t)

	// Пример ошибки при регистрации
	mockDB.On("CreateUser", newUser).Return(errors.New("user already exists"))

	req, _ = http.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"Test User", "login":"testuser", "password":"testpass"}`))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	mockDB.AssertExpectations(t)
}
