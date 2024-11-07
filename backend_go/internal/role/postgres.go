package role

import (
	"github.com/VladislavSCV/internal/models"
)

type roleHandlerDB struct {
}

func (rh roleHandlerDB) GetAll() ([]*models.Role, error) {
	return nil, nil
}

func (rh roleHandlerDB) GetById(id int) (*models.Role, error) {
	return nil, nil
}

func (rh roleHandlerDB) Create(role *models.Role) error {
	return nil
}

func (rh roleHandlerDB) Update(role *models.Role) error {
	return nil
}

func (rh roleHandlerDB) Delete(id int) error {
	return nil
}

func NewRoleHandler() RolePostgresRepository {
	return roleHandlerDB{}
}
