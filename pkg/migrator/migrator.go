package migrator

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"strings"

	"github.com/mushtaqx/mi/pkg/utils"
)

var db *sql.DB

// Migration contract
type Migration interface {
	Up()
	Down()
}

var migrations []Migration
var directory string

// Register migrations
func Register(dbConn *sql.DB, dir string, m ...Migration) {
	db = dbConn
	directory = dir
	migrations = m
	initialTable()
}

// Up : Setup database
func Up() {
	count := eachNewMigration(func(name string, Up func()) {
		Up()
		store(name)
	})
	if count == 0 {
		fmt.Println("No new migrations available")
	}
}

// Down : Take down database
func Down() {
	var count int
	for _, migration := range migrations {
		name := migrationFileName(migration)
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
func Create(name string) {
	// transform name to snake_case
	filename := utils.SnakeCase(name)
	// prefix filename with formated datetime
	filename = fmt.Sprintf("%d_%s", utils.NowSpecial(), filename)
	// concatinate a clean path for file write
	filePath := path.Clean(fmt.Sprintf("%s/%s.go", directory, filename))
	// Write stub to migration file, on error log & exit
	err := ioutil.WriteFile(filePath, []byte(stub(name)), 0774)
	fatalOnErr(err)
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

// On each new migration in migrations, runs callback with aditional data
func eachNewMigration(callback func(name string, Up func())) (count int) {
	for _, migration := range migrations {
		name := migrationFileName(migration)
		// if migration doesn't exists in DB
		if found := exists(name); !found {
			defer fmt.Printf("%s successfully migrated\n", name)
			callback(name, migration.Up)
			count++
		}
	}

	return count
}

// Get migration filename
func migrationFileName(migration Migration) string {
	// Reflect migration and convert name to snake_case
	name := utils.SnakeCase(reflect.TypeOf(migration).Elem().Name())
	// check given migration file exists
	info := utils.FileNameLike(name, directory)
	// get file name, excluding file extension
	name = strings.TrimSuffix(info.Name(), ".go")

	return name
}

// returns migration setup for name
func stub(name string) string {
	str := "package migrations\n\n"
	str += "type %s struct{}\n\n"
	str += "func (m %[1]v) Up() {\n\t// setup db\n}\n\n"
	str += "func (m %[1]v) Down() {\n\t// take down db\n}"

	return fmt.Sprintf(str, utils.UpperFirst(name))
}

// Erro check
func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
