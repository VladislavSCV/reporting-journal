package users

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/pkg"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"log"
)

type userHandlerDB struct {
	dbAndTx models.Execer
	logger  zap.Logger
}

// Структура для передачи результата
type Result struct {
	Users []models.User
	Err   error
}

func (uhp *userHandlerDB) GetUsers() ([]models.User, error) {
	rows, err := uhp.dbAndTx.Query(`
        SELECT u.id, u.first_name, u.middle_name, u.last_name, u.role_id, u.group_id, u.login, u.password, u.salt,
               r.value AS role, 
               COALESCE(g.name, 'Не указана группа') AS group_name  -- Используем COALESCE для замены NULL на строку
FROM users u
LEFT JOIN roles r ON u.role_id = r.id
LEFT JOIN groups g ON u.group_id = g.id;

    `)
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(
			&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.RoleID, &user.GroupID,
			&user.Login, &user.Hash, &user.Salt, &user.Role, &user.Group,
		); err != nil {
			return nil, pkg.LogWriteFileReturnError(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (uhp *userHandlerDB) GetStudents() ([]models.User, error) {
	var users []models.User

	query := `
        SELECT 
            u.id, 
            u.first_name, 
            u.middle_name, 
            u.last_name, 
            u.role_id, 
            u.group_id, 
            a.status,
            u.login, 
            COALESCE(r.value, 'Неизвестная роль') AS role_name, 
            COALESCE(g.name, 'Не указана группа') AS group_name
        FROM users u
        LEFT JOIN roles r ON u.role_id = r.id
        LEFT JOIN groups g ON u.group_id = g.id
        LEFT JOIN attendance a ON u.id = a.student_id
        WHERE u.role_id = 1;`

	rows, err := uhp.dbAndTx.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		var roleName, groupName string

		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.MiddleName,
			&user.LastName,
			&user.RoleID,
			&user.GroupID,
			&user.Status,
			&user.Login,
			&roleName,
			&groupName,
		)
		if err != nil {
			return nil, fmt.Errorf("ошибка при чтении строки: %w", err)
		}

		user.Role = roleName
		user.Group = &groupName

		users = append(users, user)
		log.Printf("%+v\n", user)
	}

	// Проверяем ошибки, если произошли при итерировании строк
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации строк: %w", err)
	}

	return users, nil
}

func (uhp *userHandlerDB) GetTeachers() ([]models.User, error) {
	var users []models.User
	rows, err := uhp.dbAndTx.Query("SELECT id, first_name, middle_name, last_name, role_id, group_id, login FROM users WHERE role_id = 2")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.RoleID, &user.GroupID, &user.Login)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		log.Printf("%+v\n", user)
	}
	return users, nil
}

func (uhp *userHandlerDB) UpdateToken(id int, token string) error {
	_, err := uhp.dbAndTx.Exec(`UPDATE users SET token = $1 WHERE id = $2`, token, id)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	return nil
}

// GetUserByLogin возвращает пользователя по его логину
//
//	@param login string - логин пользователя
//
//	@return models.User - пользователь
//	@return error - ошибка, если она возникла
func (uhp *userHandlerDB) GetUserByLogin(login string) (models.User, error) {
	var user models.User
	row := uhp.dbAndTx.QueryRow(`SELECT id, first_name, middle_name, last_name, role_id, group_id, login, password, salt FROM users WHERE login = $1`, login)
	err := row.Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.RoleID, &user.GroupID, &user.Login, &user.Hash, &user.Salt)
	if errors.Is(err, sql.ErrNoRows) {
		errMsg := fmt.Errorf("user %s not found", login)
		pkg.LogWriteFileReturnError(errMsg)
		return models.User{}, errMsg
	}
	if err != nil {
		errMsg := fmt.Errorf("failed to get user %s from PostgreSQL: %w", login, err)
		pkg.LogWriteFileReturnError(errMsg)
		return models.User{}, errMsg
	}

	successMsg := fmt.Sprintf("successfully retrieved user %s from PostgreSQL", login)
	pkg.LogWriteFileReturnError(fmt.Errorf(successMsg))
	return user, nil
}

func (uhp *userHandlerDB) GetUsersByGroupID(groupID int) ([]models.User, error) {
	rows, err := uhp.dbAndTx.Query(`SELECT 
            u.id, 
            u.first_name, 
            u.middle_name, 
            u.last_name, 
            u.role_id, 
            u.group_id, 
            u.login, 
            COALESCE(r.value, 'Неизвестная роль') AS role_name, 
            COALESCE(g.name, 'Не указана группа') AS group_name
        FROM users u
        LEFT JOIN roles r ON u.role_id = r.id
        LEFT JOIN groups g ON u.group_id = g.id
        WHERE u.group_id = $1;
`, groupID)
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		var roleName, groupName string

		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.MiddleName,
			&user.LastName,
			&user.RoleID,
			&user.GroupID,
			&user.Login,
			&roleName,
			&groupName,
		)
		if err != nil {
			return nil, fmt.Errorf("ошибка при чтении строки: %w", err)
		}

		user.Role = roleName
		user.Group = &groupName

		users = append(users, user)
		log.Printf("%+v\n", user)
	}
	return users, nil
}

func (uhp *userHandlerDB) GetUsersByRoleID(roleID int) ([]models.User, error) {
	rows, err := uhp.dbAndTx.Query(`SELECT id, first_name, middle_name, last_name, role_id, group_id, login, password FROM users WHERE role_id = $1`, roleID)
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.RoleID, &user.GroupID, &user.Login, &user.Hash)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (uhp *userHandlerDB) GetUserByToken(token string) (models.User, error) {
	var user models.User
	row := uhp.dbAndTx.QueryRow(`SELECT u.first_name, u.middle_name, u.last_name, r.value as role FROM users as u JOIN roles as r ON u.role_id = r.id WHERE u.token = $1`, token)
	err := row.Scan(&user.FirstName, &user.MiddleName, &user.LastName, &user.Role)
	if err == sql.ErrNoRows {
		return models.User{}, pkg.LogWriteFileReturnError(errors.New("User is not found"))
	} else if err != nil {
		return models.User{}, pkg.LogWriteFileReturnError(err)
	}
	return user, nil
}

// GetUserById возвращает пользователя по его ID
//
//	@param id int - ID пользователя
//
//	@return models.User - пользователь
//	@return error - ошибка, если она возникла
//
// GetUserById возвращает пользователя по его ID
func (uhp *userHandlerDB) GetUserById(id int) (models.User, error) {
	var user models.User
	row := uhp.dbAndTx.QueryRow(`SELECT u.first_name, u.middle_name, u.last_name, r.value, g.name, login, password FROM users as u LEFT JOIN roles as r ON u.role_id = r.id LEFT JOIN groups as g ON u.group_id = g.id WHERE u.id = $1`, id)
	err := row.Scan(&user.FirstName, &user.MiddleName, &user.LastName, &user.Role, &user.Group, &user.Login, &user.Hash)
	if err == sql.ErrNoRows {
		return models.User{}, pkg.LogWriteFileReturnError(errors.New("User is not found"))
	} else if err != nil {
		return models.User{}, pkg.LogWriteFileReturnError(err)
	}
	user.ID = id
	fmt.Println(user)
	return user, nil
}

// CreateUser создает нового пользователя
//
//	@param user *models.User - пользователь, который будет создан
//
//	@return error - ошибка, если она возникла
func (uhp *userHandlerDB) CreateStudent(user *models.User) (models.User, string, error) {
	log.Println("STUDENT")
	var responseUser models.User
	// Генерация соли и хеш пароля
	hashResult, err := pkg.CreateHashWithSalt(user.Hash)
	if err != nil {
		return models.User{}, "", pkg.LogWriteFileReturnError(err)
	}

	// Logging
	pkg.LogWriteFileReturnError(fmt.Errorf("successfully generated hash with salt for user %s", user.Login))

	// Проверяем, существует ли уже пользователь с таким логином
	var count int
	err = uhp.dbAndTx.QueryRow(`SELECT COUNT(*) FROM users WHERE login = $1`, user.Login).Scan(&count)
	if err != nil {
		return models.User{}, "", pkg.LogWriteFileReturnError(err)
	}

	// Logging
	pkg.LogWriteFileReturnError(fmt.Errorf("successfully checked if user %s already exists", user.Login))

	if count > 0 {
		// Если такой логин уже существует, возвращаем ошибку
		return models.User{}, "", errors.New("пользователь с таким логином уже существует")
	}

	// Logging
	pkg.LogWriteFileReturnError(fmt.Errorf("user %s does not exist, so creating new user", user.Login))

	// Сохранение пользователя с солью и хешем пароля
	var insertId int
	err = uhp.dbAndTx.QueryRow(`INSERT INTO users (first_name, middle_name, last_name, role_id, group_id, login, password, salt) 
                               VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, first_name, middle_name, last_name, role_id, group_id, login`,
		user.FirstName, user.MiddleName, user.LastName, user.RoleID, user.GroupID, user.Login, hashResult.Hash, hashResult.Salt).
		Scan(&responseUser.ID, &responseUser.FirstName, &responseUser.MiddleName, &responseUser.LastName, &responseUser.RoleID, &responseUser.GroupID, &responseUser.Login)
	if err != nil {
		return models.User{}, "", pkg.LogWriteFileReturnError(err)
	}

	// Генерация токена
	token, err := pkg.GenerateJWT(insertId, user.RoleID)
	if err != nil {
		uhp.logger.Error("failed to generate token",
			zap.Int("id", insertId),
			zap.Error(err),
		)
		return models.User{}, "", pkg.LogWriteFileReturnError(errors.New("failed to generate token"))
	}

	// Logging
	pkg.LogWriteFileReturnError(fmt.Errorf("successfully created user %s", user.Login))

	return responseUser, token, nil
}

func (uhp *userHandlerDB) CreateTeacher(user *models.User) (models.User, string, error) {
	log.Println("TEACHER")
	var responseUser models.User
	// Генерация соли и хеш пароля
	hashResult, err := pkg.CreateHashWithSalt(user.Hash)
	if err != nil {
		return models.User{}, "", pkg.LogWriteFileReturnError(err)
	}

	// Logging
	pkg.LogWriteFileReturnError(fmt.Errorf("successfully generated hash with salt for user %s", user.Login))

	// Проверяем, существует ли уже пользователь с таким логином
	var count int
	err = uhp.dbAndTx.QueryRow(`SELECT COUNT(*) FROM users WHERE login = $1`, user.Login).Scan(&count)
	if err != nil {
		return models.User{}, "", pkg.LogWriteFileReturnError(err)
	}

	// Logging
	pkg.LogWriteFileReturnError(fmt.Errorf("successfully checked if user %s already exists", user.Login))

	if count > 0 {
		// Если такой логин уже существует, возвращаем ошибку
		return models.User{}, "", errors.New("пользователь с таким логином уже существует")
	}

	// Logging
	pkg.LogWriteFileReturnError(fmt.Errorf("user %s does not exist, so creating new user", user.Login))

	// Сохранение пользователя с солью и хешем пароля
	var insertId int
	err = uhp.dbAndTx.QueryRow(`INSERT INTO users (first_name, middle_name, last_name, role_id, group_id, login, password, salt) 
                               VALUES ($1, $2, $3, $4, NULL, $5, $6, $7) RETURNING id, first_name, middle_name, last_name, role_id, group_id, login`,
		user.FirstName, user.MiddleName, user.LastName, user.RoleID, user.Login, hashResult.Hash, hashResult.Salt).
		Scan(&responseUser.ID, &responseUser.FirstName, &responseUser.MiddleName, &responseUser.LastName, &responseUser.RoleID, &responseUser.GroupID, &responseUser.Login)
	if err != nil {
		return models.User{}, "", pkg.LogWriteFileReturnError(err)
	}

	// Генерация токена
	token, err := pkg.GenerateJWT(insertId, user.RoleID)
	if err != nil {
		uhp.logger.Error("failed to generate token",
			zap.Int("id", insertId),
			zap.Error(err),
		)
		return models.User{}, "", pkg.LogWriteFileReturnError(errors.New("failed to generate token"))
	}

	// Logging
	pkg.LogWriteFileReturnError(fmt.Errorf("successfully created user %s", user.Login))

	return responseUser, token, nil
}

func (uhp *userHandlerDB) GetCuratorGroups(id int) (models.User, error) {
	query := `
		SELECT 
			g.id AS group_id,
			g.name AS group_name
		FROM teacher_groups tg
		INNER JOIN groups g ON tg.group_id = g.id
		WHERE tg.teacher_id = $1;
	`

	var user models.User
	err := uhp.dbAndTx.QueryRow(query, id).Scan(&user.GroupID, &user.Group)
	if err != nil {
		return models.User{}, pkg.LogWriteFileReturnError(err)
	}

	return user, nil
}

// UpdateUser обновляет существующего пользователя
//
//	@param id int - ID пользователя, который будет обновлен
//	@param updates map[string]string - поля, которые будут обновлены
//
//	@return error - ошибка, если она возникла
func (uhp *userHandlerDB) UpdateUser(id int, updates map[string]string) error {
	// Проверка наличия пользователя
	_, err := uhp.GetUserById(id)
	if err != nil {
		return pkg.LogWriteFileReturnError(err) // Здесь лучше уточнить ошибку
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
	if uhp == nil {
		return fmt.Errorf("nil pointer to userHandlerDB")
	}

	//uhp.logger.Info("attempting to delete user", zap.Int("user_id", id))
	_, err := uhp.dbAndTx.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		//uhp.logger.Error("failed to delete user", zap.Int("user_id", id), zap.Error(err))
		return pkg.LogWriteFileReturnError(err)
	}
	//uhp.logger.Info("successfully deleted user", zap.Int("user_id", id))
	return nil
}

// ConnToDB возвращает соединение с базой данных PostgreSQL
//
//	@param connStr string - строка, содержащая информацию о подключении к базе
//
//	@return *sql.DB - готовое соединение с базой
//
//	@error error - ошибка, если она возникла
func ConnToDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	return db, nil
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
	db, err := ConnToDB(connStr)
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}
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
