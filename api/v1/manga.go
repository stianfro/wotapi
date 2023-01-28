package v1

import (
	"encoding/json"
	"net/http"

	// go-sqlite3 is needed to initialize the database
	_ "github.com/mattn/go-sqlite3"

	"github.com/rs/zerolog/log"
	"github.com/stianfro/chi-no-wadachi/models"
	"github.com/stianfro/chi-no-wadachi/utils"
)

// CreateManga godoc
// @Summary Create a manga
// @Description Creates a manga in the database
// @Tags manga
// @Accept json
// @Produce json
// @Param manga body models.Manga true "Manga"
// @Success 201 {object} models.Manga
// @Router /api/v1/manga [post]
func CreateManga(w http.ResponseWriter, r *http.Request) {
	var manga models.Manga

	err := json.NewDecoder(r.Body).Decode(&manga)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to decode JSON body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uuid, err := utils.NewUUID()
	if err != nil {
		log.Error().
			Msg("Failed to generate UUID")
		w.WriteHeader(http.StatusInternalServerError)
	}
	manga.ID = uuid

	log.Info().
		Str("id", manga.ID).
		Str("title", manga.Title).
		Str("author", manga.Author).
		Str("publisher", manga.Publisher).
		Str("magazine", manga.Magazine).
		Msg("Manga created")

	db, err := utils.InitDB()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to initialize database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare(`INSERT INTO manga (id, title, author, magazine, publisher) VALUES (?,?,?,?,?)`)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to prepare statement")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(manga.ID, manga.Title, manga.Author, manga.Magazine, manga.Publisher)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to execute statement")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
