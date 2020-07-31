# MiGo
### Simple database migrator in Go.

Usage
Create new migration
```go
import mi "github.com/mushtaqx/migo/migrator"

// ...

// Create a new migration
func Create(name string) {
    mi.Create(name, "data/migrations")
}

```


Register created migrations so they can be migrated

```go
import mi "github.com/mushtaqx/migo/migrator"

// ...

// Register takes *sql.DB as first arg
mi.Register(
    db,
    new(CreateExampleTable),
)
```

```go
import mi "github.com/mushtaqx/migo/migrator"

// ...

mi.Up()
mi.Down()
```
