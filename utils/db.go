package utils

import (
	"database/sql"
	"os"

	// Import the postgres driver
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// InitDB initializes the database
func InitDB() (*sqlx.DB, error) {
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	dbDriver := os.Getenv("DB_DRIVER")

	conn, err := sql.Open(dbDriver, dbConnectionString)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to open database")
		return nil, err
	}

	db := sqlx.NewDb(conn, dbDriver)

	return db, nil
}
