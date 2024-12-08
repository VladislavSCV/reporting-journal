package schedules

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/pkg"
	"log"
)

type scheduleHandlerDB struct {
	dbAndTx models.Execer
}

func (s *scheduleHandlerDB) GetSchedulers() ([]models.Schedule, error) {
	var schedules []models.Schedule
	rows, err := s.dbAndTx.Query("SELECT s.id AS schedule_id, g.name AS group_name, s.day_of_week, sub.name AS subject_name, " +
		"CONCAT(u.first_name, ' ', u.middle_name, ' ', COALESCE(u.last_name, '')) AS teacher_name, s.location FROM schedules s " +
		"JOIN groups g ON s.group_id = g.id JOIN subjects sub ON s.subject_id = sub.id " +
		"JOIN users u ON s.teacher_id = u.id ORDER BY s.day_of_week, s.id;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schedule models.Schedule
		err := rows.Scan(&schedule.ScheduleID, &schedule.GroupName, &schedule.DayOfWeek, &schedule.SubjectName, &schedule.TeacherName, &schedule.Location)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (s *scheduleHandlerDB) GetScheduleForGroup(id int) ([]models.Schedule, error) {
	var schedules []models.Schedule
	rows, err := s.dbAndTx.Query("SELECT s.id AS schedule_id, g.name AS group_name, s.day_of_week, sub.name AS subject_name, "+
		"CONCAT(u.first_name, ' ', u.middle_name, ' ', COALESCE(u.last_name, '')) AS teacher_name, s.location "+
		"FROM schedules s JOIN groups g ON s.group_id = g.id JOIN subjects sub ON s.subject_id = sub.id "+
		"JOIN users u ON s.teacher_id = u.id WHERE g.id = $1 ORDER BY s.day_of_week, s.id;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schedule models.Schedule
		err := rows.Scan(&schedule.ScheduleID, &schedule.GroupName, &schedule.DayOfWeek, &schedule.SubjectName, &schedule.TeacherName, &schedule.Location)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return schedules, nil
}

func (s *scheduleHandlerDB) CreateSchedule(schedule models.CreateSchedule) error {
	//var createSchedule models.CreateSchedule
	log.Println(schedule)
	_, err := s.dbAndTx.Exec("INSERT INTO schedules (group_id, day_of_week, subject_id, teacher_id, location) VALUES ($1, $2, $3, $4, $5)", schedule.GroupID, schedule.DayOfWeek, schedule.SubjectID, schedule.TeacherID, schedule.Location)
	if err != nil {
		return err
	}
	return nil
}

func (s *scheduleHandlerDB) UpdateSchedule(id int, updatedSchedule map[string]interface{}) error {
	_, err := s.GetScheduleForGroup(id)
	if err != nil {
		return err
	}
	log.Println(updatedSchedule)

	if len(updatedSchedule) == 0 {
		return pkg.LogWriteFileReturnError(errors.New("Обновлений нет"))
	}

	// Формирование запроса
	query := "UPDATE schedules SET "
	var args []interface{}
	i := 1

	for k, v := range updatedSchedule {
		query += fmt.Sprintf("%s = $%d, ", k, i) // Используем $ для подстановки
		args = append(args, v)
		i++
	}

	query = query[:len(query)-2]
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	log.Println("Executing query:", query, args)

	_, err = s.dbAndTx.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (s *scheduleHandlerDB) DeleteSchedule(id int) error {
	_, err := s.dbAndTx.Exec("DELETE FROM schedules WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

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
func checkConPostgres(dbConn *sql.DB) error {
	if err := dbConn.Ping(); err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	return nil
}

// NewUserPostgresHandlerDB возвращает UserHandlerDB, готовый к работе с БД
//
//	@param connStr string - строка, содержащая информацию о подключении к базе
//
//	@return UserHandlerDB - готовый UserHandlerDB
func NewSchedulePostgresHandlerDB(connStr string) SchedulePostgresRepository {
	db, err := ConnToDB(connStr)
	if err != nil {
		return nil
	}
	if err := db.Ping(); err != nil {
		return nil
	}
	return &scheduleHandlerDB{dbAndTx: db}
}
