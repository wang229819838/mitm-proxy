// pkg/db.go

package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase(path string) (*sql.DB, error) {
	return sql.Open("sqlite3", path)
}

func ValidateUser(db *sql.DB, username, password string) bool {
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&storedPassword)
	if err != nil {
		return false
	}
	return password == storedPassword
}
