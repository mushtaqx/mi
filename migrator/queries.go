package migrator

// Create initial "migrations" table for storing migrations
func initTable() {
	q := "CREATE TABLE IF NOT EXISTS migrations ( name varchar(255) NOT NULL );"
	_, err := db.Exec(q)
	fatalOnErr(err)
}

// Check migration existence
func exists(name string) bool {
	var res string
	query := "SELECT name FROM migrations WHERE name = ? LIMIT 1"
	err := db.QueryRow(query, name).Scan(&res)

	return err == nil
}

// Store migration
func store(name string) {
	query := "INSERT INTO migrations(name) VALUES(?)"
	_, err := db.Query(query, name)
	fatalOnErr(err)
}

// Destroy migration
func destroy(name string) {
	query := "DELETE FROM migrations WHERE name=?"
	_, err := db.Query(query, name)
	fatalOnErr(err)
}
