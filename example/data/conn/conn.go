package conn

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

// DB connection
func DB() *sql.DB {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), dataSourceName())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func dataSourceName() string {
	return fmt.Sprintf("%s:%s@/%s?parseTime=true&charset=utf8",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
}
