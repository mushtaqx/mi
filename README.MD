# Mi
### Simple database migrator in Go.
Mi is inspired by Laravels migrator altough that one is more coupled with the framework, this merely provides with a way to easly create some migration files where you can specify Up and Down operations.

#### *When migration Up() runs, its filename gets stored in DB to not run it again.*

#### *By calling Down you will destroy all migration records in migrations table*

#### Usage

Create migration
```go
import mi "github.com/mushtaqx/mi/pkg/migrator"

// ...

// Create migration
mi.Create("createUsersTable")
mi.Create("CreateRolesTable")
```


Register migrations

```go

import (
    "github.com/mushtaqx/mi/internal/data/conn"
    mi "github.com/mushtaqx/mi/pkg/migrator"
)

// ...

mi.Register(
    // Provide a connection (*sql.DB)
    conn.DB(),
    // migrations directory
    "internal/data/migrations",
    // Migrations to run
    new(CreateUsersTable),
    new(CreateRolesTable),
)
```

```go
import mi "github.com/mushtaqx/mi/pkg/migrator"

// ...

// Run Up for all migrations
mi.Up()
// Run Down for all migrations
mi.Down()
```
