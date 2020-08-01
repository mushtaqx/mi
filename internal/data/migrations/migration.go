package migrations

import (
	"github.com/mushtaqx/mi/internal/data/conn"
	mi "github.com/mushtaqx/mi/pkg/migrator"
)

func init() {
	// Example registration
	mi.Register(
		// Provide connection
		conn.DB(),
		// migrations directory
		"internal/data/migrations",
		// Migrations to run
		new(CreateUsersTable),
		new(CreateRolesTable),
	)
}

// Up : Setup database
func Up() {
	mi.Up()
}

// Down : Take down database
func Down() {
	mi.Down()
}

// Create a new migration
func Create(name string) {
	mi.Create(name)
}
