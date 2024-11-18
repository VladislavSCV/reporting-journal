package note

import (
	"database/sql"
	"fmt"
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/pkg"
)

type noteHandlerDB struct {
	dbAndTx models.Execer
}

func (nh *noteHandlerDB) CreateNote(note models.Note) error {
	_, err := nh.dbAndTx.Exec("INSERT INTO notes (title, body, group_id, user_id) VALUES ($1, $2, $3, $4)", note.Title, note.Body, note.GroupId, note.UserId)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}

	return nil
}

func (nh *noteHandlerDB) GetNote(id int) (models.Note, error) {
	var note models.Note
	nh.dbAndTx.QueryRow("SELECT title, body FROM notes WHERE id = $1", id).Scan(&note.Title, &note.Body)
	return note, nil
}

func (nh *noteHandlerDB) GetNotes() ([]models.Note, error) {
	var notes []models.Note
	rows, err := nh.dbAndTx.Query("SELECT title, body FROM notes")
	if err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	defer rows.Close()
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.Title, &note.Body); err != nil {
			return nil, pkg.LogWriteFileReturnError(err)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, pkg.LogWriteFileReturnError(err)
	}
	return notes, nil
}

func (nh *noteHandlerDB) UpdateNote(id int, newNote map[string]interface{}) error {
	_, err := nh.GetNote(id)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	if len(newNote) == 0 {
		return nil
	}

	query := "UPDATE notes SET "
	var args []interface{}
	i := 1
	for k, v := range newNote {
		query += fmt.Sprint("%s = $%d, ", k, v)
		args = append(args, id)
		i++
	}

	query = query[:len(query)-2]
	query += fmt.Sprintf("WHERE id = $%d", id)
	args = append(args, id)

	_, err = nh.dbAndTx.Exec(query, args...)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
	}
	return nil
}

func (nh *noteHandlerDB) DeleteNote(noteId int) error {
	_, err := nh.dbAndTx.Exec("DELETE FROM notes WHERE id = $1", noteId)
	if err != nil {
		return pkg.LogWriteFileReturnError(err)
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

func NewNotePostgresHandlerDB(connStr string) NotePostgresRepository {
	db, err := ConnToDB(connStr)
	if err != nil {
		pkg.LogWriteFileReturnError(err)
	}
	checkConPostgres(db)
	return &noteHandlerDB{dbAndTx: db}
}
