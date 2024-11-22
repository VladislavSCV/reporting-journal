package subjects

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/pkg"
)

type subjectHandlerDB struct {
	dbAndTx models.Execer
}

func (s *subjectHandlerDB) GetSubjects() ([]models.Subject, error) {
	var subjects []models.Subject
	rows, err := s.dbAndTx.Query("SELECT * FROM subjects")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var subject models.Subject
		err := rows.Scan(&subject.Id, &subject.Name)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}
	return subjects, nil
}

func (s *subjectHandlerDB) GetSubjectById(id int) (models.Subject, error) {
	var subject models.Subject
	err := s.dbAndTx.QueryRow("SELECT * FROM subjects WHERE id = $1", id).Scan(&subject.Id, &subject.Name)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Subject{}, err
	} else if err != nil {
		return models.Subject{}, err
	}

	return subject, nil
}

func (s *subjectHandlerDB) CreateSubject(subject models.Subject) error {
	_, err := s.dbAndTx.Exec("INSERT INTO subjects (name) VALUES ($1)", subject.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *subjectHandlerDB) UpdateSubject(id int, UpdatedSubjects map[string]interface{}) error {
	// Проверка наличия пользователя
	_, err := s.GetSubjectById(id)
	if err != nil {
		return pkg.LogWriteFileReturnError(err) // Здесь лучше уточнить ошибку
	}

	// Проверка наличия полей для обновления
	if len(UpdatedSubjects) == 0 {
		return fmt.Errorf("no fields to update for user with ID %d", id)
	}

	// Формирование запроса
	query := "UPDATE subjects SET "
	var args []interface{}
	i := 1

	for k, v := range UpdatedSubjects {
		query += fmt.Sprintf("%s = $%d, ", k, i) // Используем $ для подстановки
		args = append(args, v)
		i++
	}

	// Удаляем последнюю запятую
	query = query[:len(query)-2]
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	// Выполняем запрос
	_, err = s.dbAndTx.Exec(query, args...)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}

	return nil
}

func (s *subjectHandlerDB) DeleteSubject(id int) error {
	_, err := s.dbAndTx.Exec("DELETE FROM subjects WHERE id = $1", id)
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

func checkConPostgres(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func NewSubjectPostgresHandlerDB(connStr string) SubjectPostgresRepository {
	db, err := ConnToDB(connStr)
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}
	checkConPostgres(db)
	return &subjectHandlerDB{dbAndTx: db}
}
