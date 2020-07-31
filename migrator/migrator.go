package migrator

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/mushtaqx/migo/utils"
)

var db *sql.DB

// Migration contract
type Migration interface {
	Up()
	Down()
}

var migrations []Migration

// Register migrations
func Register(dbConn *sql.DB, m ...Migration) {
	db = dbConn
	migrations = m
	initialTable()
}

// Up : Setup database
func Up() {
	var count int
	for _, migration := range migrations {
		name := getFilename(migration)
		// if migration doesn't exists in DB,run T.Up() and store migration
		if found := exists(name); !found {
			defer fmt.Printf("%s successfully migrated\n", name)
			migration.Up()
			store(name) // store to db
			count++
		}
	}
	if count == 0 {
		fmt.Println("No new migrations available")
	}
}

// Down : Take down database
func Down() {
	var count int
	for _, migration := range migrations {
		name := getFilename(migration)
		if found := exists(name); found {
			defer fmt.Printf("%s successfully dropped\n", name)
			migration.Down()
			destroy(name)
			count++
		}
	}
	if count == 0 {
		fmt.Println("Migrations table empty")
	}
}

// Create migration file
func Create(name, location string) {
	// transform name to snake_case
	filename := utils.SnakeCase(name)
	// prefix filename with formated datetime
	filename = strings.Replace(time.Now().Format("020106030405.000"), ".", "", 1) + "_" + filename
	// concatinate a clean path for file write
	filePath := path.Clean(fmt.Sprintf("%s/%s.go", location, filename))
	// Write stub to migration file, on error log & exit
	fatalOnErr(ioutil.WriteFile(filePath, []byte(stub(name)), 0774))
	fmt.Printf("%s successfully created", filename)
}

// create table for storing migrations
func initialTable() {
	q := "CREATE TABLE IF NOT EXISTS migrations ( name varchar(255) NOT NULL );"
	_, err := db.Exec(q)
	fatalOnErr(err)
}

// Store migration name
func store(name string) {
	query := "INSERT INTO migrations(name) VALUES(?)"
	_, err := db.Query(query, name)
	fatalOnErr(err)
}

// Destroy migration by name
func destroy(name string) {
	query := "DELETE FROM migrations WHERE name=?"
	_, err := db.Query(query, name)
	fatalOnErr(err)
}

// Check migration by name
func exists(name string) bool {
	var res string
	query := "SELECT name FROM migrations WHERE name = ? LIMIT 1"
	err := db.QueryRow(query, name).Scan(&res)

	return err == nil
}

// Get migration filename
func getFilename(migration Migration) string {
	// Reflect migration and convert name to snake_case
	name := utils.SnakeCase(reflect.TypeOf(migration).Elem().Name())
	// check given migration file exists
	info := utils.FileExist("data/migrations", name)
	// get file name, excluding file extension
	name = strings.TrimSuffix(info.Name(), ".go")

	return name
}

// Read migration stub
func stub(name string) string {
	stub, err := ioutil.ReadFile(path.Clean("migrator/__stub__"))
	fatalOnErr(err)
	if unicode.IsUpper(rune(name[0])) {
		return fmt.Sprintf(string(stub), name)
	}
	// transform lowercase name to uppercase
	upper := strings.ToUpper(string(name[0]))
	upper += name[1:]

	return fmt.Sprintf(string(stub), upper)
}

// Erro check
func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
