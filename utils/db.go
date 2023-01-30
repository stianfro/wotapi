package utils

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// InitDB initializes the database
func InitDB() (*sqlx.DB, error) {
	conn, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to open database")
		return nil, err
	}

	db := sqlx.NewDb(conn, "sqlite3")

	return db, nil
}
