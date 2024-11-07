package role

import (
	"github.com/VladislavSCV/internal/models"
)

type RolePostgresRepository interface {
	GetAll() ([]*models.Role, error)
	GetById(id int) (*models.Role, error)
	Create(role *models.Role) error
	Update(role *models.Role) error
	Delete(id int) error
}

type RoleApiRepository interface {
	GetAll() ([]models.Role, error)
	GetById(id int) (models.Role, error)
	Create(role *models.Role) error
	Update(role *models.Role) error
	Delete(id int) error
}
