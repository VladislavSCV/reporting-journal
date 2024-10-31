package users

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/VladislavSCV/internal/model"
	"github.com/VladislavSCV/pkg"
	_ "github.com/lib/pq"
)

type userHandlerDB struct {
	dbAndTx model.Execer
}

// GetUsers возвращает список всех пользователей
//
//	@return []model.User - список пользователей
//	@return error - ошибка, если она возникла
func (uhp *userHandlerDB) GetUsers() (*[]model.User, error) {
	rows, err := uhp.dbAndTx.Query("SELECT * FROM users")
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
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
	return &users, nil
}

// GetUserByLogin возвращает пользователя по его логину
//
//	@param login string - логин пользователя
//
//	@return model.User - пользователь
//	@return error - ошибка, если она возникла
func (uhp *userHandlerDB) GetUserByLogin(login string) (model.User, error) {
	var user model.User
	row := uhp.dbAndTx.QueryRow(`SELECT id, name, role_id, group_id, login, password from users WHERE login = $1`, login)
	err := row.Scan(&user.ID, &user.Name, &user.RoleID, &user.GroupID, &user.Login, &user.Password)
	if err != nil {
		return model.User{}, pkg.LogWriteFileReturnError(err)
	}
	return user, nil
}

// GetUserById возвращает пользователя по его ID
//
//	@param id int - ID пользователя
//
//	@return model.User - пользователь
//	@return error - ошибка, если она возникла
func (uhp *userHandlerDB) GetUserById(id int) (model.User, error) {
	var user model.User
	row := uhp.dbAndTx.QueryRow(`SELECT name, role_id, group_id, login, password FROM users WHERE id = $1`, id)

	err := row.Scan(&user.Name, &user.RoleID, &user.GroupID, &user.Login, &user.Password)
	if err != nil {
		return model.User{}, pkg.LogWriteFileReturnError(err)
	}
	return user, nil
}

// CreateUser создает нового пользователя
//
//	@param user *model.User - пользователь, который будет создан
//
//	@return error - ошибка, если она возникла
func (uhp *userHandlerDB) CreateUser(user *model.User) error {
	password, err := pkg.GenerateFromPassword(user.Password)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	fmt.Println(user.Password)

	_, err = uhp.dbAndTx.Exec(`INSERT INTO users (name, role_id, group_id, login, password) VALUES ($1, $2, $3, $4, $5)`, user.Name, user.RoleID, user.GroupID, user.Login, password)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}

	return nil
}

// UpdateUser обновляет существующего пользователя
//
//	@param id int - ID пользователя, который будет обновлен
//	@param updates map[string]string - поля, которые будут обновлены
//
//	@return error - ошибка, если она возникла
func (uhp *userHandlerDB) UpdateUser(StrId string, updates map[string]string) error {
	// Преобразование строки в число
	id, err := strconv.Atoi(StrId)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}

	// Проверка наличия пользователя
	_, err = uhp.GetUserById(id)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}

	// Проверка наличия полей для обновления
	if len(updates) == 0 {
		return fmt.Errorf("no fields to update for user with ID %d", id)
	}

	// Формирование запроса
	query := "UPDATE users SET "
	var args []interface{}
	i := 1

	for k, v := range updates {
		query += fmt.Sprintf("%s = $%d, ", k, i) // Используем $ для подстановки
		args = append(args, v)
		i++
	}

	// Удаляем последнюю запятую
	query = query[:len(query)-2]
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	// Выполняем запрос
	_, err = uhp.dbAndTx.Exec(query, args...)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}

	return nil
}

// DeleteUser удаляет существующего пользователя
//
//	@param id int - ID пользователя, который будет удален
//
//	@return error - ошибка, если она возникла
func (uhp *userHandlerDB) DeleteUser(id int) error {
	_, err := uhp.dbAndTx.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	return nil
}

// ConnToDB возвращает соединение с базой данных PostgreSQL
//
//	@param connStr string - строка, содержащая информацию о подключении к базе
//
//	@return *sql.DB - готовое соединение с базой
//
//	@error error - ошибка, если она возникла
func ConnToDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	pkg.LogWriteFileReturnError(err)
	return db
}

// CheckConn проверяет соединение с базой данных PostgreSQL
//
//	@param db *sql.DB - соединение с базой данных
//
//	@return error - ошибка, если она возникла
func checkConPostgres(dbConn *sql.DB) {
	pkg.LogWriteFileReturnError(dbConn.Ping())
}

// NewUserPostgresHandlerDB возвращает UserHandlerDB, готовый к работе с БД
//
//	@param connStr string - строка, содержащая информацию о подключении к базе
//
//	@return UserHandlerDB - готовый UserHandlerDB
func NewUserPostgresHandlerDB(connStr string) UserPostgresRepository {
	db := ConnToDB(connStr)
	checkConPostgres(db)
	return &userHandlerDB{dbAndTx: db}
}

// NewUserPostgresHandlerDBWithoutConnStr возвращает готовый UserHandlerDB с готовым соединением
//
//	@param db *sql.DB - готовое соединение с базой данных
//
//	@return UserHandlerDB - готовый UserHandlerDB
func NewUserPostgresHandlerDBWithoutConnStr(db *sql.DB) UserPostgresRepository {
	//db := ConnToDB(connStr)
	checkConPostgres(db)
	return &userHandlerDB{dbAndTx: db}
}
