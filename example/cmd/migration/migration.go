package migration

import (
	"fmt"
	"os"

	"github.com/mushtaqx/mi/example/data/conn"
	. "github.com/mushtaqx/mi/example/data/migrations"
	"github.com/mushtaqx/mi/migrator"
)

func init() {
	// Initial Mi setup
	migrator.Setup(
		conn.DB(),                 // DB connection
		"example/data/migrations", // migrations directory
	)
	// Register migrations to run
	migrator.Register(
		new(CreateUsersTable),
		new(CreateRolesTable),
	)
}

// Execute example command
func Execute() {
	if len(os.Args) > 2 && os.Args[1] == "mi" {
		handle(os.Args)
	}
}

func handle(args []string) {
	switch args[2] {
	case "new":
		migrator.New(args[3])
	case "up":
		migrator.Up()
	case "down":
		migrator.Down()
	case "redo":
		migrator.Redo()
	case "status":
		migrator.Status()
	default:
		fmt.Printf("Command \"%s\" not found\n", args[1])
	}
}
