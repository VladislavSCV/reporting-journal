package postgres

import (
	"database/sql"

	"github.com/VladislavSCV/internal/model"
	"github.com/VladislavSCV/pkg"
	_ "github.com/lib/pq"
)

type UserHandlerDB interface {
	GetUsers() ([]model.User, error)
	GetUserByLogin(login string) (model.User, error)
	GetUserById(id int) (model.User, error)
	CreateUser(user model.User) error
	UpdateUser(id int, updates map[string]string) error
	DeleteUser(id int) error
}

type userHandlerDB struct {
	db *sql.DB
}

// GetUsers возвращает список всех пользователей.
//
//	GET /api/v1/users
//
//	Responses:
//	  200 OK
//	  500 Internal Server Error
func (uh *userHandlerDB) GetUsers() ([]model.User, error) {
	rows, err := uh.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.RoleID, &user.GroupID, &user.Login, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (uh *userHandlerDB) GetUserByLogin(login string) (model.User, error) {
	//rows, err := uh.db.Query(`SELECT id, name, role_id, group_id, login, password`)
	return model.User{}, nil
}

func (uh *userHandlerDB) GetUserById(id int) (model.User, error) {
	return model.User{}, nil
}

func (uh *userHandlerDB) CreateUser(user model.User) error {
	return nil
}

func (uh *userHandlerDB) UpdateUser(id int, updates map[string]string) error {
	return nil
}

func (uh *userHandlerDB) DeleteUser(id int) error {
	return nil
}

// ConnToDB - функция, которая открывает соединение с базой данных Postgres
// по заданной строке подключения.
func ConnToDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	pkg.CError(err)
	return db
}

// CheckConn - функция, которая проверяет соединение с базой данных.
//
//	Если соединение не работает, то она вызывает pkg.CError с ошибкой
//	соединения.
func CheckConn(db *sql.DB) {
	pkg.CError(db.Ping())
}

// NewUserPostgresHandlerDB - конструктор для UserHandlerDB, который принимает
// строку подключения к базе данных в формате Postgres.
//
// Он создает подключение к базе данных, проверяет его, и возвращает указатель
// на UserHandlerDB, хранящий это подключение.
func NewUserPostgresHandlerDB(connStr string) UserHandlerDB {
	db := ConnToDB(connStr)
	CheckConn(db)
	return &userHandlerDB{db: db}
}

// NewUserPostgresHandlerDBWithoutConnStr - конструктор для UserHandlerDB, принимает уже
// существующий соединениe с базой данных.
//
//	NewUserPostgresHandlerDBWithoutConnStr - это вспомогательная функция, которая
//	принимает уже существующий объект sql.DB, проверяет его на работоспособность
//	и возвращает объект UserHandlerDB.
func NewUserPostgresHandlerDBWithoutConnStr(db *sql.DB) UserHandlerDB {
	//db := ConnToDB(connStr)
	CheckConn(db)
	return &userHandlerDB{db: db}
}
