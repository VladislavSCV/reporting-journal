package role

import (
	"database/sql"
	"errors"
	"log"

	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/pkg"
)

type roleHandlerDB struct {
	dbAndTx models.Execer
}

func (rh roleHandlerDB) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	q, err := rh.dbAndTx.Query("SELECT id, value FROM roles")
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	for q.Next() {
		var role models.Role
		if err := q.Scan(&role.ID, &role.Value); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (rh roleHandlerDB) GetRole(id int) (*models.Role, error) {
	var role models.Role
	err := rh.dbAndTx.QueryRow("SELECT * FROM roles WHERE id = $1", id).Scan(&role.ID, &role.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.LogWriteFileReturnError(errors.New("role is not found"))
		}
		return nil, pkg.LogWriteFileReturnError(err)
	}

	return &role, nil

}

func (rh roleHandlerDB) CreateRole(role *models.Role) (*models.Role, error) {
	var respRole models.Role
	err := rh.dbAndTx.QueryRow("INSERT INTO roles (value) VALUES ($1) RETURNING id, value", role.Value).Scan(&respRole.ID, &respRole.Value)
	if err != nil {
		return &models.Role{}, pkg.LogWriteFileReturnError(err)
	}
	log.Println(respRole)
	return &respRole, nil
}

func (rh roleHandlerDB) UpdateRole(role *models.Role) error {
	q, err := rh.dbAndTx.Exec("UPDATE roles SET value = $1 WHERE id = $2", role.Value, role.ID)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}

	_, err = q.RowsAffected()
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	return nil
}

func (rh roleHandlerDB) DeleteRole(id int) error {
	q, err := rh.dbAndTx.Exec("DELETE FROM roles WHERE id = $1", id)

	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}

	_, err = q.RowsAffected()
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	return nil
}

func checkConPostgres(db *sql.DB) {
	if db == nil {
		panic("db is nil")
	}
	p := db.Ping()
	if p != nil {
		panic(p)
	}

}

func ConnToDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	return db, nil
}

func NewRolePostgresHandler(connStr string) RolePostgresRepository {
	db, err := ConnToDB(connStr)
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}
	checkConPostgres(db)
	return &roleHandlerDB{dbAndTx: db}
}
