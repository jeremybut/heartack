package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"

	"fmt"
	"log"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	fmt.Println("Initializing database")
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrate",
	}

	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
