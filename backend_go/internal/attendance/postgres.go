package attendance

import (
	"database/sql"
	"fmt"
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/pkg"
	"time"
)

type attendanceHandler struct {
	dbAndTx models.Execer
}

func (ah *attendanceHandler) UpdateAttendance(studentID int, status string) error {
	dateTime := time.Now()
	attendanceDate := dateTime.Format("2006-01-02")

	//// Используем транзакцию
	//tx, err := ah.dbAndTx.Begin()
	//if err != nil {
	//	return fmt.Errorf("error starting transaction: %v", err)
	//}
	//defer tx.Rollback()

	// Обновляем посещаемость
	query := `
		UPDATE attendance
		SET status = $1, updated_at = CURRENT_TIMESTAMP
		WHERE student_id = $2 AND date = $3
	`
	result, err := ah.dbAndTx.Exec(query, status, studentID, attendanceDate)
	if err != nil {
		return fmt.Errorf("error updating attendance: %v", err)
	}

	// Проверяем, был ли затронут хотя бы один ряд
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no attendance record found for student %d on date %s", studentID, attendanceDate)
	}

	//// Завершаем транзакцию
	//if err := tx.Commit(); err != nil {
	//	return fmt.Errorf("error committing transaction: %v", err)
	//}

	return nil
}

func (ah *attendanceHandler) AddAttendance(studentID int, status string) error {
	dateTime := time.Now()
	attendanceDate := dateTime.Format("2006-01-02")

	// Проверяем, есть ли уже запись
	var count int
	checkQuery := `
		SELECT COUNT(*) FROM attendance WHERE student_id = $1 AND date = $2
	`
	err := ah.dbAndTx.QueryRow(checkQuery, studentID, attendanceDate).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking attendance existence: %v", err)
	}

	if count > 0 {
		return fmt.Errorf("attendance record already exists for student %d on date %s", studentID, attendanceDate)
	}

	// Добавляем новую запись
	query := `
		INSERT INTO attendance (student_id, date, status)
		VALUES ($1, $2, $3)
	`
	_, err = ah.dbAndTx.Exec(query, studentID, attendanceDate, status)
	if err != nil {
		return fmt.Errorf("error adding attendance: %v", err)
	}

	return nil
}

func (ah *attendanceHandler) GetAttendanceForStudent(studentID int) ([]Attendance, error) {
	query := `
		SELECT id, student_id, date, status, created_at, updated_at
		FROM attendance
		WHERE student_id = $1
	`
	rows, err := ah.dbAndTx.Query(query, studentID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving attendance: %v", err)
	}
	defer rows.Close()

	var attendanceRecords []Attendance
	for rows.Next() {
		var record Attendance
		if err := rows.Scan(&record.ID, &record.StudentID, &record.Date, &record.Status, &record.CreatedAt, &record.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning attendance record: %v", err)
		}
		attendanceRecords = append(attendanceRecords, record)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return attendanceRecords, nil
}

func (ah *attendanceHandler) DeleteAttendance(studentID int, attendanceDate string) error {
	query := `
		DELETE FROM attendance
		WHERE student_id = $1 AND date = $2
	`
	_, err := ah.dbAndTx.Exec(query, studentID, attendanceDate)
	if err != nil {
		return fmt.Errorf("error deleting attendance: %v", err)
	}
	return nil
}

//
//func main() {
//	// Пример добавления записи о посещаемости
//	err := addAttendance(1, "2024-11-26", "Прогул")
//	if err != nil {
//		log.Fatalf("Error adding attendance: %v", err)
//	}
//
//	// Пример получения записей для студента
//	attendanceRecords, err := getAttendanceForStudent(1)
//	if err != nil {
//		log.Fatalf("Error getting attendance: %v", err)
//	}
//	fmt.Println("Attendance records:", attendanceRecords)
//
//	// Пример обновления записи о посещаемости
//	err = updateAttendance(1, "2024-11-26", "Болеет")
//	if err != nil {
//		log.Fatalf("Error updating attendance: %v", err)
//	}
//
//	// Пример удаления записи о посещаемости
//	err = deleteAttendance(1, "2024-11-26")
//	if err != nil {
//		log.Fatalf("Error deleting attendance: %v", err)
//	}
//}

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
func NewAttendancePostgresHandlerDB(connStr string) AttendancePostgresRepository {
	db, err := ConnToDB(connStr)
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}
	checkConPostgres(db)
	return &attendanceHandler{dbAndTx: db}
}
