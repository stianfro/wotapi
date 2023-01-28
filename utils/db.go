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

	if _, err := db.Exec("SELECT * FROM manga"); err != nil {
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
	}

	if _, err := db.Exec("SELECT * FROM mangaVolumes"); err != nil {
		_, err = db.Exec(`
	  CREATE TABLE mangaVolumes (
		  id TEXT PRIMARY KEY,
			mangaID TEXT NOT NULL,
			number INTEGER NOT NULL,
			title TEXT NOT NULL,
			releaseDate TEXT NOT NULL,
			isbn TEXT NOT NULL,
			FOREIGN KEY (mangaID) REFERENCES manga(id)
			)
		`)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Failed to create mangaVolumes table")
			return nil, err
		}
	}

	if _, err := db.Exec("SELECT * FROM mangaChapters"); err != nil {
		_, err = db.Exec(`
	  CREATE TABLE mangaChapters (
		  id TEXT PRIMARY KEY,
			volumeID TEXT NOT NULL,
			number INTEGER NOT NULL,
			title TEXT NOT NULL,
			FOREIGN KEY (volumeID) REFERENCES mangaVolumes(id)
			)
		`)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Failed to create mangaChapters table")
			return nil, err
		}
	}

	return db, nil
}
