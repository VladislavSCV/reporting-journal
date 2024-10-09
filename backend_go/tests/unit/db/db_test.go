package db

import (
	"regexp"
	"testing"

	"backend_go/internal/database/postgres"
	"backend_go/internal/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создаем ожидаемые строки в ответе
	rows := sqlmock.NewRows([]string{"id", "name", "role_id", "group_id", "login", "password"}).
		AddRow(1, "John", 1, 1, "john123", "password").
		AddRow(2, "Jane", 2, 2, "jane123", "password")

	// Ожидаем запрос к базе данных
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users")).
		WillReturnRows(rows)

	// Инициализируем наш хендлер с мок базой данных
	handler := postgres.NewUserPostgresHandlerDB(db)

	// Вызываем метод GetUsers
	users, err := handler.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, []model.User{
		{ID: 1, Name: "John", RoleID: 1, GroupID: 1, Login: "john123", Password: "password"},
		{ID: 2, Name: "Jane", RoleID: 2, GroupID: 2, Login: "jane123", Password: "password"},
	}, users)

	// Проверяем, что все ожидания были выполнены
	assert.NoError(t, mock.ExpectationsWereMet())
}
