package database

import (
	"database/sql"
	"os"

	"go.uber.org/zap"
)

var (
	Db     *sql.DB
	err    error
	logger *zap.Logger
)

func InitDB() {
	connStr := os.Getenv("POSTGRES_URL")
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("Error connecting to database: ", zap.Error(err))
	}
}

func CreateTables() {
	_, err = Db.Exec(`
		CREATE TABLE IF NOT EXISTS notes (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL
		)
	`)
	if err != nil {
		logger.Fatal("Error creating table: ", zap.Error(err))
	}
}
