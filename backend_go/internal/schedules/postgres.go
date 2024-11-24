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
	rows, err := s.dbAndTx.Query("SELECT * FROM schedules")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schedule models.Schedule
		err := rows.Scan(&schedule.ID, &schedule.GroupID, &schedule.DayOfWeek, &schedule.StartTime, &schedule.EndTime, &schedule.SubjectID, &schedule.TeacherID, &schedule.Location)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (s *scheduleHandlerDB) GetScheduleById(id int) (models.Schedule, error) {
	var schedule models.Schedule
	err := s.dbAndTx.QueryRow("SELECT * FROM schedules WHERE id = $1", id).Scan(&schedule.ID, &schedule.GroupID, &schedule.DayOfWeek, &schedule.StartTime, &schedule.EndTime, &schedule.SubjectID, &schedule.TeacherID, &schedule.Location)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Schedule{}, err
	} else if err != nil {
		return models.Schedule{}, err
	}

	return schedule, nil
}

func (s *scheduleHandlerDB) CreateSchedule(schedule models.Schedule) error {
	_, err := s.dbAndTx.Exec("INSERT INTO schedules (group_id, day_of_week, start_time, end_time, subject_id, teacher_id, location) VALUES ($1, $2, $3, $4, $5, $6, $7)", schedule.GroupID, schedule.DayOfWeek, schedule.StartTime, schedule.EndTime, schedule.SubjectID, schedule.TeacherID, schedule.Location)
	if err != nil {
		return err
	}
	return nil
}

func (s *scheduleHandlerDB) UpdateSchedule(id int, updatedSchedule map[string]interface{}) error {
	_, err := s.GetScheduleById(id)
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
func checkConPostgres(dbConn *sql.DB) {
	pkg.LogWriteFileReturnError(dbConn.Ping())
}

// NewUserPostgresHandlerDB возвращает UserHandlerDB, готовый к работе с БД
//
//	@param connStr string - строка, содержащая информацию о подключении к базе
//
//	@return UserHandlerDB - готовый UserHandlerDB
func NewSchedulePostgresHandlerDB(connStr string) SchedulePostgresRepository {
	db, err := ConnToDB(connStr)
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}
	checkConPostgres(db)
	return &scheduleHandlerDB{dbAndTx: db}
}
