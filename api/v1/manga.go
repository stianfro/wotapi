package v1

import (
	"encoding/json"
	"net/http"

	// go-sqlite3 is needed to initialize the database
	_ "github.com/mattn/go-sqlite3"

	"github.com/rs/zerolog/log"
	"github.com/stianfro/wotapi/models"
	"github.com/stianfro/wotapi/services"
)

// HTTPHandler is the HTTP handler for the API
type HTTPHandler struct {
	service *services.Service
}

// NewHTTPHandler creates a new HTTPHandler with the given service
func NewHTTPHandler(s *services.Service) *HTTPHandler {
	return &HTTPHandler{service: s}
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), status)
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
// @Failure 500 {object} string
// @Router /manga/{id} [get]
func (h *HTTPHandler) GetManga(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/v1/manga/"):]

	manga, err := h.service.GetManga(id)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to get manga")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, manga, http.StatusOK)
}

// ListManga godoc
// @Summary List all manga
// @Description Lists all manga in the database
// @Tags manga
// @Produce json
// @Success 200 {array} models.Manga
// @Failure 500 {object} string
// @Router /manga [get]
func (h *HTTPHandler) ListManga(w http.ResponseWriter, r *http.Request) {
	manga, err := h.service.ListManga()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to list manga")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, manga, http.StatusOK)
}

// CreateManga godoc
// @Summary Create a manga
// @Description Creates a manga in the database
// @Tags manga
// @Accept json
// @Produce json
// @Param manga body models.Manga true "Manga"
// @Success 201 {object} models.Manga
// @Failure 500 {object} string
// @Router /manga [post]
func (h *HTTPHandler) CreateManga(w http.ResponseWriter, r *http.Request) {
	data := &models.Manga{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to decode JSON body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	manga, err := h.service.CreateManga(data)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to create manga")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, manga, http.StatusCreated)
}

// GetVolume godoc
// @Summary Get a volume by ID
// @Description Gets a volume from the database
// @Tags volume
// @Produce json
// @Param id path string true "Volume ID"
// @Success 200 {object} models.Volume
// @Failure 500 {object} string
// @Router /manga/volume/{id} [get]
func (h *HTTPHandler) GetVolume(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/v1/manga/volume/"):]

	volume, err := h.service.GetVolume(id)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to get volume")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, volume, http.StatusOK)
}

// ListVolumes godoc
// @Summary List all volumes
// @Description Lists all volumes in the database
// @Tags volume
// @Produce json
// @Success 200 {array} models.Volume
// @Failure 500 {object} string
// @Router /manga/volume [get]
func (h *HTTPHandler) ListVolume(w http.ResponseWriter, r *http.Request) {
	volumes, err := h.service.ListVolumes()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to list volumes")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, volumes, http.StatusOK)
}

// CreateVolume godoc
// @Summary Create a volume
// @Description Creates a volume in the database
// @Tags volume
// @Accept json
// @Produce json
// @Param volume body models.Volume true "Volume"
// @Success 201 {object} models.Volume
// @Failure 500 {object} string
// @Router /manga/volume [post]
func (h *HTTPHandler) CreateVolume(w http.ResponseWriter, r *http.Request) {
	data := &models.Volume{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to decode JSON body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	volume, err := h.service.CreateVolume(data)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to create volume")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, volume, http.StatusCreated)
}
