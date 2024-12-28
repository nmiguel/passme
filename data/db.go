package data

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var customDBPath string

// SetCustomDBPath overrides the default database path (for tests, etc.).
func SetCustomDBPath(path string) {
	customDBPath = path
}

// getDBPath ensures the correct folder exists and returns the sqlite database path.
func getDBPath() (string, error) {
	if customDBPath != "" {
		return customDBPath, nil
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(configDir, "passme")
	if err := os.MkdirAll(appDir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(appDir, "passme.db"), nil
}

// openDB opens the database and ensures the table is created.
func openDB() (*sql.DB, error) {
	dbPath, err := getDBPath()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	if err := initDB(db); err != nil {
		return nil, err
	}
	return db, nil
}

// initDB creates the table if it doesn't exist.
func initDB(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS keys (
		alias TEXT PRIMARY KEY,
		token TEXT NOT NULL
	);
	`
	_, err := db.Exec(query)
	return err
}

// InsertKey adds a new alias/token pair to the database.
func InsertKey(alias, token string) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT OR REPLACE INTO keys (alias, token) VALUES (?, ?);"
	_, err = db.Exec(query, alias, token)
	return err
}

// GetAllKeys returns all alias/token pairs from the database.
func GetAllKeys() ([]Key, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT alias, token FROM keys;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keys []Key
	for rows.Next() {
		var k Key
		if err := rows.Scan(&k.Alias, &k.Token); err != nil {
			return nil, err
		}
		keys = append(keys, k)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return keys, nil
}

// DeleteKey removes an alias/token entry from the database.
func DeleteKey(alias string) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM keys WHERE alias = ?;", alias)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no entry found for alias: %s", alias)
	}
	return nil
}
