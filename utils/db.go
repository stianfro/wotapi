package utils

import (
	"database/sql"

	"github.com/rs/zerolog/log"
)

// InitDB initializes the database and creates the manga table.
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to open database")
		return nil, err
	}

	if _, err := db.Exec("SELECT * FROM manga"); err == nil {
		// Table already exists
		return db, nil
	}

	_, err = db.Exec(`
  	CREATE TABLE manga (
  		id TEXT PRIMARY KEY,
  		title TEXT NOT NULL,
  		author TEXT NOT NULL,
  		magazine TEXT NOT NULL,
  		publisher TEXT NOT NULL
  	)
	`)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to create manga table")
		return nil, err
	}

	return db, nil
}
