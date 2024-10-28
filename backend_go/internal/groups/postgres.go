package groups

import (
	"database/sql"

	"github.com/VladislavSCV/internal/model"
	"github.com/VladislavSCV/pkg"
)

type groupHandlerDB struct {
	db *sql.DB
}

// NewGroupRepository создает новый репозиторий для работы с группами
func NewGroupRepository(db *sql.DB) GroupPostgresRepository {
	return &groupHandlerDB{db: db}
}

// CreateGroup добавляет новую группу в БД
func (gh *groupHandlerDB) CreateGroup(group *model.Group) error {
	_, err := gh.db.Exec("INSERT INTO groups (name) VALUES ($1)", group.Name)
	return pkg.LogWriteFileReturnError(err)
}

// GetGroupByID получает группу по ID
func (gh *groupHandlerDB) GetGroupByID(id int) (*model.Group, error) {
	group := &model.Group{}
	err := gh.db.QueryRow("SELECT id, name FROM groups WHERE id = $1", id).Scan(&group.Id, &group.Name)
	if err == sql.ErrNoRows {
		return nil, nil // если группа не найдена, возвращаем nil
	}
	return group, pkg.LogWriteFileReturnError(err)
}

// GetAllGroups возвращает список всех групп
func (gh *groupHandlerDB) GetAllGroups() ([]*model.Group, error) {
	rows, err := gh.db.Query("SELECT id, name FROM groups")
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var groups []*model.Group
	for rows.Next() {
		group := &model.Group{}
		if err := rows.Scan(&group.Id, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

// UpdateGroup обновляет данные группы
func (gh *groupHandlerDB) UpdateGroup(group *model.Group) error {
	_, err := gh.db.Exec("UPDATE groups SET name = $1 WHERE id = $2", group.Name, group.Id)
	return pkg.LogWriteFileReturnError(err)
}

// DeleteGroup удаляет группу по ID
func (gh *groupHandlerDB) DeleteGroup(id int) error {
	_, err := gh.db.Exec("DELETE FROM groups WHERE id = $1", id)
	return pkg.LogWriteFileReturnError(err)
}

// AddStudentToGroup добавляет студента в группу, обновляя его group_id
func (gh *groupHandlerDB) AddStudentToGroup(studentID, groupID int) error {
	_, err := gh.db.Exec("UPDATE users SET group_id = $1 WHERE id = $2", groupID, studentID)
	return pkg.LogWriteFileReturnError(err)
}

// RemoveStudentFromGroup удаляет студента из группы, обновляя его group_id на NULL
func (gh *groupHandlerDB) RemoveStudentFromGroup(studentID int) error {
	_, err := gh.db.Exec("UPDATE users SET group_id = NULL WHERE id = $1", studentID)
	return pkg.LogWriteFileReturnError(err)
}

// GetStudentsByGroupID возвращает список студентов, принадлежащих указанной группе
func (gh *groupHandlerDB) GetStudentsByGroupID(groupID int) ([]*model.User, error) {
	rows, err := gh.db.Query("SELECT id, name, role_id, group_id, login FROM users WHERE group_id = $1", groupID)
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var students []*model.User
	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.RoleID, &user.GroupID, &user.Login); err != nil {
			return nil, err
		}
		students = append(students, user)
	}
	return students, nil
}

// FindGroupsByName находит группы по имени (или его части)
func (gh *groupHandlerDB) FindGroupsByName(name string) ([]*model.Group, error) {
	rows, err := gh.db.Query("SELECT id, name FROM groups WHERE name ILIKE '%' || $1 || '%'", name)
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var groups []*model.Group
	for rows.Next() {
		group := &model.Group{}
		if err := rows.Scan(&group.Id, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
