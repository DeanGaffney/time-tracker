package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "time-tracker.db")

	if err != nil {
		log.Fatal("Failed to load database time-tracker.db")
	}

	return db
}

func CloseConnection(db *sql.DB) {
	db.Close()
}

func ExecuteStatement(db *sql.DB, statement string, args ...any) {

	if len(args) > 0 {

		_, err := db.Exec(statement, args)
		if err != nil {
			log.Printf("%q: %s\n", err, statement)
		}
	} else {

		_, err := db.Exec(statement)
		if err != nil {
			log.Printf("%q: %s\n", err, statement)
		}
	}
}

func SetUpTables(db *sql.DB) {

	createTableStatement := `
  CREATE TABLE IF NOT EXISTS time_tracking (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    start_date TEXT,
    end_date TEXT,
    description TEXT
  )
  `

	ExecuteStatement(db, createTableStatement)

}
