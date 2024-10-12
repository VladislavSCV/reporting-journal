package unit

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VladislavSCV/internal/database/postgres"
	"github.com/VladislavSCV/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (name, role_id, group_id, login, password) VALUES ($1, $2, $3, $4, $5)")).
		WithArgs("John", 1, 1, "john123", "password").
		WillReturnResult(sqlmock.NewResult(1, 1))

	handler := postgres.NewUserPostgresHandlerDBWithoutConnStr(db)

	user := model.User{
		Name:     "John",
		RoleID:   1,
		GroupID:  1,
		Login:    "john123",
		Password: "password",
	}

	err = handler.CreateUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Настраиваем мок-ответ для выборки пользователя
	rows := sqlmock.NewRows([]string{"id", "name", "role_id", "group_id", "login", "password"}).
		AddRow(1, "John", 1, 1, "john123", "password")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE login=$1")).
		WithArgs("john123").
		WillReturnRows(rows)

	handler := postgres.NewUserPostgresHandlerDBWithoutConnStr(db)

	// Тестируем метод
	user, err := handler.GetUserByLogin("john123")
	assert.NoError(t, err)
	assert.Equal(t, model.User{
		ID:       1,
		Name:     "John",
		RoleID:   1,
		GroupID:  1,
		Login:    "john123",
		Password: "password",
	}, user)

	assert.NoError(t, mock.ExpectationsWereMet())
}
