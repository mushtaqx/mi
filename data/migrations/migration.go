package migrations

import (
	. "github.com/mushtaqx/migo/data/conn"
	mi "github.com/mushtaqx/migo/migrator"
)

func init() {
	mi.Register(
		DB(),
		// new(CreateExampleTable),
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
	mi.Create(name, "data/migrations")
}
