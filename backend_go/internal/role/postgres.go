package role

import (
	"github.com/VladislavSCV/internal/models"
)

type roleHandlerDB struct {
}

func (rh roleHandlerDB) GetRoles() ([]*models.Role, error) {
	return nil, nil
}

func (rh roleHandlerDB) GetRole(id int) (*models.Role, error) {
	return nil, nil
}

func (rh roleHandlerDB) CreateRole(role *models.Role) error {
	return nil
}

func (rh roleHandlerDB) UpdateRole(role *models.Role) error {
	return nil
}

func (rh roleHandlerDB) DeleteRole(id int) error {
	return nil
}

func NewRoleHandler() RolePostgresRepository {
	return roleHandlerDB{}
}
