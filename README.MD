# Mi
Easly create migration files where you can specify Up and Down operations in GO.

#### Usage

```go

import (
    "github.com/mushtaqx/mi/example/data/conn"
    "github.com/mushtaqx/mi/migrator"
)

// Initial Mi setup
migrator.Setup(
    conn.DB(),                  // DB connection
    "example/data/migrations", // migrations directory
)

// Create new migration
migrator.New("createUsersTable")

// Register migrations to run
migrator.Register(
    new(CreateUsersTable),
    new(CreateRolesTable),
)

// Runs Up for all registered migrations
migrator.Up()

// Runs Down for all registered migrations
migrator.Down()

// Runs Down and then Up for all registered migrations
migrator.Redo()

// Check the status for applied migrations
migrator.Status()
```
