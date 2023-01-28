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
// @Router /manga [post]
func CreateManga(w http.ResponseWriter, r *http.Request) {
	var manga models.Manga

	w.Header().Set("Content-Type", "application/json")

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

// ListManga godoc
// @Summary List all manga
// @Description Lists all manga in the database
// @Tags manga
// @Produce json
// @Success 200 {array} models.Manga
// @Router /manga [get]
func ListManga(w http.ResponseWriter, r *http.Request) {
	var manga []models.Manga

	w.Header().Set("Content-Type", "application/json")

	db, err := utils.InitDB()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to initialize database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT * FROM manga")
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to query database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var m models.Manga
		err = rows.Scan(&m.ID, &m.Title, &m.Author, &m.Magazine, &m.Publisher)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Failed to scan row")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		manga = append(manga, m)
	}

	err = json.NewEncoder(w).Encode(manga)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to encode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetManga godoc
// @Summary Get a manga
// @Description Gets a manga from the database
// @Tags manga
// @Produce json
// @Param id path string true "Manga ID"
// @Success 200 {object} models.Manga
// @Router /manga/{id} [get]
func GetManga(w http.ResponseWriter, r *http.Request) {
	var manga models.Manga

	w.Header().Set("Content-Type", "application/json")

	db, err := utils.InitDB()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to initialize database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("SELECT * FROM manga WHERE id=?")
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to prepare statement")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := r.URL.Path[len("/api/v1/manga/"):]
	err = stmt.QueryRow(id).Scan(&manga.ID, &manga.Title, &manga.Author, &manga.Magazine, &manga.Publisher)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to query database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(manga)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to encode JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
