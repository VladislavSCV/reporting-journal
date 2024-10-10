package postgres

import (
	"database/sql"

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
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (uh *userHandlerDB) GetUserByLogin(login string) (model.User, error) {
}

func (uh *userHandlerDB) GetUserById(id int) (model.User, error) {
}

func (uh *userHandlerDB) CreateUser(user model.User) error {
}

func (uh *userHandlerDB) UpdateUser(id int, updates map[string]string) error {
	return nil
}

func (uh *userHandlerDB) DeleteUser(id int) error {
}

func ConnToDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	pkg.CError(err)
	return db
}

	pkg.CError(db.Ping())
}


}
