package groups

import (
	"database/sql"
	"testing"

	"github.com/VladislavSCV/internal/models"
	_ "github.com/lib/pq"
)

func setupTestDB() (*sql.DB, *sql.Tx, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=31415926 dbname=testing sslmode=disable")
	if err != nil {
		return nil, nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		db.Close()
		return nil, nil, err
	}

	return db, tx, nil
}

func TestCreateGroup(t *testing.T) {
	var group models.Group
	group.Name = "22ИС3-2"

	// Подключаемся к тестовой базе данных
	db, tx, err := setupTestDB()
	if err != nil {
		t.Fatalf("Ошибка настройки тестовой базы данных: %v", err)
	}
	defer db.Close()
	defer tx.Rollback() // Откат транзакции после теста

	// Создаем обработчик группы с доступом к транзакции
	groupHandler := groupHandlerDB{dbAndTx: tx}

	// Вызываем метод создания группы
	err = groupHandler.CreateGroup(&group)
	if err != nil {
		t.Errorf("Ошибка при создании группы: %v", err)
	}

	// Проверяем, что группа была добавлена с использованием транзакции `tx`
	var countGroup int
	err = tx.QueryRow("SELECT COUNT(*) FROM groups WHERE name = $1", group.Name).Scan(&countGroup)
	if err != nil {
		t.Fatalf("Ошибка проверки данных: %v", err)
	}

	// Проверяем, что количество групп соответствует ожиданию
	if countGroup != 1 {
		t.Errorf("Ожидалось 1, получено %d", countGroup)
	}
}
