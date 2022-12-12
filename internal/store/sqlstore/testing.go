package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

// TestDB...
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()
	// Открытие подключения к БД
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	// Проверка доступности БД
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	// Очистка БД
	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
		}

		db.Close()
	}
}
