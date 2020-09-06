package query

import (
	"database/sql"
	"log"
)

var db *sql.DB

// Setup initial "migrations" table for storing migrations
func Setup(DB *sql.DB) {
	db = DB
	q := "CREATE TABLE IF NOT EXISTS migrations ( name VARCHAR(255) NOT NULL, applied_at DATETIME DEFAULT CURRENT_TIMESTAMP );"
	_, err := db.Exec(q)
	fatalOnErr(err)
}

// All : Get all migrations
func All() (rows *sql.Rows) {
	query := "SELECT * FROM migrations"
	rows, err := db.Query(query)
	fatalOnErr(err)

	return rows
}

// Exists : Check migration existence
func Exists(name string) bool {
	var res string
	query := "SELECT name FROM migrations WHERE name = ? LIMIT 1"
	err := db.QueryRow(query, name).Scan(&res)

	return err == nil
}

// Store migration
func Store(name string) {
	query := "INSERT INTO migrations(name) VALUES(?)"
	_, err := db.Query(query, name)
	fatalOnErr(err)
}

// Destroy migration
func Destroy(name string) {
	query := "DELETE FROM migrations WHERE name=?"
	_, err := db.Query(query, name)
	fatalOnErr(err)
}

// Erro check
func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
