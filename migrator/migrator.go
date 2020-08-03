package migrator

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/mushtaqx/mi/pkg/utils"
)

var db *sql.DB

// Migration contract
type Migration interface {
	Up() bool
	Down() bool
}

var (
	directory  string
	migrations []Migration
	// print messages
	messageCreated         = "New\t%s\n"
	messageUp              = "UP\t%s\n"
	messageDown            = "DOWN\t%s\n"
	messageExists          = "EXISTS\t%s\n"
	messageNotAvailable    = "No new migrations available"
	messageTableEmpty      = "Migrations table empty"
	messageDirNotSpecified = "No migrations directory found, register it in Register func."
)

// Setup migrator
func Setup(dbConn *sql.DB, dir string) {
	db, directory = dbConn, dir
	initTable()
}

// Register migrations
func Register(m ...Migration) {
	migrations = m
}

// New migration file
func New(name string) {
	if directory == "" {
		log.Fatal(messageDirNotSpecified)
	}
	// transform name to snake_case
	filename := utils.SnakeCase(name)
	// if migration file exists, return
	if migration, _ := utils.FileNameLike(filename, directory); migration != nil {
		fmt.Fprintf(os.Stderr, messageExists, migration.Name())
		return
	}
	// prefix filename with formated datetime
	filename = fmt.Sprintf("%d_%s", utils.NowSpecial(), filename)
	// concatinate a clean path for file write
	filePath := path.Clean(fmt.Sprintf("%s/%s.go", directory, filename))
	// Write stub to migration file, on error log & exit
	err := ioutil.WriteFile(filePath, []byte(stub(name)), 0774)
	fatalOnErr(err)
	fmt.Fprintf(os.Stdout, messageCreated, filename)
}

// Up : Setup database
func Up() {
	var count int
	for _, migration := range migrations {
		n := filename(migration)
		if !exists(n) && migration.Up() {
			store(n)
			fmt.Fprintf(os.Stdout, messageUp, n)
			count++
		}
	}
	if count == 0 {
		fmt.Fprintf(os.Stderr, messageNotAvailable)
	}
}

// Down : Take down database
func Down() {
	var count int
	for _, migration := range migrations {
		n := filename(migration)
		if exists(n) && migration.Down() {
			destroy(n)
			fmt.Fprintf(os.Stdout, messageDown, n)
			count++
		}
	}
	if count == 0 {
		fmt.Fprintf(os.Stderr, messageTableEmpty)
	}
}

// Get migration scaffolding for name
func stub(name string) string {
	str := "package migrations\n\n"
	str += "type %s struct{}\n\n"
	str += "func (m %[1]v) Up() bool {\n\t// setup db\n\treturn true\n}\n\n"
	str += "func (m %[1]v) Down() bool {\n\t// take down db\n\treturn true\n}"

	return fmt.Sprintf(str, utils.ToUpperFirst(name))
}

// Get migration filename
func filename(migration Migration) string {
	// Reflect migration and convert name to snake_case
	name := utils.SnakeCase(reflect.TypeOf(migration).Elem().Name())
	// check if migration file exists in migrations directory
	file, _ := utils.FileNameLike(name, directory)
	if file == nil {
		return ""
	}
	name = strings.TrimSuffix(file.Name(), ".go")

	return name
}

// Erro check
func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
