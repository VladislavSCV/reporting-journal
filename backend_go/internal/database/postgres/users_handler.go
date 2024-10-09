package postgres

import (
	"database/sql"
	"errors"

	"backend_go/internal/model"
	"backend_go/pkg"
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
	row := uh.db.QueryRow("SELECT * FROM users WHERE login = $1", login)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.RoleID, &user.GroupID, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}
	return user, nil
}

func (uh *userHandlerDB) GetUserById(id int) (model.User, error) {
	row := uh.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.RoleID, &user.GroupID, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}
	return user, nil
}

func (uh *userHandlerDB) CreateUser(user model.User) error {
	_, err := uh.db.Exec("INSERT INTO users (name, role_id, group_id, login, password) VALUES ($1, $2, $3, $4, $5)",
		user.Name, user.RoleID, user.GroupID, user.Login, user.Password)
	return err
}

func (uh *userHandlerDB) UpdateUser(id int, updates map[string]string) error {
	//if len(updates) == 0 {
	//	return errors.New("no updates provided")
	//}
	//
	//setClause := ""
	//args := []interface{}{}
	//i := 1
	//for key, value := range updates {
	//	if setClause != "" {
	//		setClause += ", "
	//	}
	//	setClause += key + " = $" + strconv.Itoa(i)
	//	args = append(args, value)
	//	i++
	//}
	//args = append(args, id)
	//
	//query := fmt.Sprintf(`UPDATE users SET %s WHERE id = $%d`, setClause, len(args))
	//_, err := uh.db.Exec(query, args...)
	//return err
	return nil
}

func (uh *userHandlerDB) DeleteUser(id int) error {
	_, err := uh.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func ConnToDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	pkg.CError(err)
	return db
}

func checkConn(db *sql.DB) {
	pkg.CError(db.Ping())
}

func NewUserPostgresHandlerDB(conn *sql.DB) UserHandlerDB {
	checkConn(conn)

	return &userHandlerDB{db: conn}
}
