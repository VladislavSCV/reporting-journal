package groups

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/VladislavSCV/internal/model"
	"github.com/VladislavSCV/pkg"
)

type groupHandlerDB struct {
	db *sql.DB
}

func (ghp *groupHandlerDB) GetGroups() (*[]model.Group, error) {
	rows, err := ghp.db.Query("SELECT * FROM groups")
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var groups []model.Group
	for rows.Next() {
		group := model.Group{}
		err = rows.Scan(&group.Id, &group.Name)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	if rows.Err() != nil {
		return nil, pkg.LogWriteFileReturnError(rows.Err())
	}
	return &groups, nil
}

// GetStudentsByGroupName returns a list of students in a given group
//
//	@param groupName string - name of the group
//
//	@return *[]model.User - list of students
//	@return error - error, if it occurs
func (ghp *groupHandlerDB) GetStudentsByGroupName(groupName string) (*[]model.User, error) {
	groupsId := ghp.db.QueryRow("SELECT id FROM groups WHERE name = $1", groupName)

	var groupId int
	err := groupsId.Scan(&groupId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.LogWriteFileReturnError(err)
		}
		return nil, pkg.LogWriteFileReturnError(err)
	}

	rows, err := ghp.db.Query("SELECT id, name, role_id, group_id, login FROM users WHERE group_id = $1", groupId)
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()

	var students []model.User
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.RoleID, &user.GroupID, &user.Login)
		if err != nil {
			return nil, err
		}
		students = append(students, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &students, nil
}

func (ghp *groupHandlerDB) GetGroupById(id int) (model.Group, error) {
	var group model.Group
	err := ghp.db.QueryRow("SELECT name FROM groups WHERE id = $1", id).Scan(&group.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Group{}, pkg.LogWriteFileReturnError(fmt.Errorf("group not found: id %d", id))
		}
		return model.Group{}, pkg.LogWriteFileReturnError(err)
	}
	group.Id = id // Можно установить ID группы
	return group, nil
}

func (ghp *groupHandlerDB) CreateGroup(group *model.Group) error {
	_, err := ghp.db.Exec("INSERT INTO groups (name) VALUES ($1)", group.Name)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	return nil
}

func (ghp *groupHandlerDB) DeleteStudentFromGroup(studentId int) error {
	// Проверяем, существует ли студент
	var exists bool
	err := ghp.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", studentId).Scan(&exists)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	if !exists {
		return pkg.LogWriteFileReturnError(fmt.Errorf("student not found: id %d", studentId))
	}

	// Удаляем студента из группы
	_, err = ghp.db.Exec("UPDATE users SET group_id = NULL WHERE id = $1", studentId)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	return nil
}

func (ghp *groupHandlerDB) DeleteGroup(id int) error {
	_, err := ghp.db.Exec("DELETE FROM groups WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkg.LogWriteFileReturnError(fmt.Errorf("group not found: id %d", id))
		}
		return pkg.LogWriteFileReturnError(err)
	}
	return nil
}
