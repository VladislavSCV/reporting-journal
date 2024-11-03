package groups

import (
	"database/sql"

	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/pkg"
)

// groupHandlerDB структура для работы с группами.
type groupHandlerDB struct {
	dbAndTx models.Execer // Используется для выполнения запросов
}

// NewGroupRepository создает новый репозиторий для работы с группами
func NewGroupRepository(dbAndTx models.Execer) GroupPostgresRepository {
	return &groupHandlerDB{dbAndTx: dbAndTx}
}

// CreateGroup добавляет новую группу в БД
func (ghp *groupHandlerDB) CreateGroup(group *models.Group) error {
	_, err := ghp.dbAndTx.Exec("INSERT INTO groups (name) VALUES ($1)", group.Name)
	return pkg.LogWriteFileReturnError(err)
}

// GetGroupByID получает группу по ID
func (ghp *groupHandlerDB) GetGroupByID(id int) (*models.Group, error) {
	group := &models.Group{}
	err := ghp.dbAndTx.QueryRow("SELECT id, name FROM groups WHERE id = $1", id).Scan(&group.Id, &group.Name)
	if err == sql.ErrNoRows {
		return nil, nil // Если группа не найдена, возвращаем nil
	}
	return group, pkg.LogWriteFileReturnError(err)
}

// GetAllGroups возвращает список всех групп
func (ghp *groupHandlerDB) GetAllGroups() ([]*models.Group, error) {
	rows, err := ghp.dbAndTx.Query("SELECT id, name FROM groups")
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var groups []*models.Group
	for rows.Next() {
		group := &models.Group{}
		if err := rows.Scan(&group.Id, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

// UpdateGroup обновляет данные группы
func (ghp *groupHandlerDB) UpdateGroup(group *models.Group) error {
	_, err := ghp.dbAndTx.Exec("UPDATE groups SET name = $1 WHERE id = $2", group.Name, group.Id)
	return pkg.LogWriteFileReturnError(err)
}

// DeleteGroup удаляет группу по ID
func (ghp *groupHandlerDB) DeleteGroup(id int) error {
	_, err := ghp.dbAndTx.Exec("DELETE FROM groups WHERE id = $1", id)
	return pkg.LogWriteFileReturnError(err)
}

// AddStudentToGroup добавляет студента в группу, обновляя его group_id
func (ghp *groupHandlerDB) AddStudentToGroup(studentID, groupID int) error {
	_, err := ghp.dbAndTx.Exec("UPDATE users SET group_id = $1 WHERE id = $2", groupID, studentID)
	return pkg.LogWriteFileReturnError(err)
}

// RemoveStudentFromGroup удаляет студента из группы, обновляя его group_id на NULL
func (ghp *groupHandlerDB) RemoveStudentFromGroup(studentID int) error {
	_, err := ghp.dbAndTx.Exec("UPDATE users SET group_id = NULL WHERE id = $1", studentID)
	return pkg.LogWriteFileReturnError(err)
}

// GetStudentsByGroupID возвращает список студентов, принадлежащих указанной группе
func (ghp *groupHandlerDB) GetStudentsByGroupID(groupID int) ([]*models.User, error) {
	rows, err := ghp.dbAndTx.Query("SELECT id, name, role_id, group_id, login FROM users WHERE group_id = $1", groupID)
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var students []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.RoleID, &user.GroupID, &user.Login); err != nil {
			return nil, err
		}
		students = append(students, user)
	}
	return students, nil
}

// FindGroupsByName находит группы по имени (или его части)
func (ghp *groupHandlerDB) FindGroupsByName(name string) ([]*models.Group, error) {
	rows, err := ghp.dbAndTx.Query("SELECT id, name FROM groups WHERE name ILIKE '%' || $1 || '%'", name)
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var groups []*models.Group
	for rows.Next() {
		group := &models.Group{}
		if err := rows.Scan(&group.Id, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
