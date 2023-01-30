package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/stianfro/wotapi/utils"
)

// Manga is a struct that represents a manga
type Manga struct {
	// ID is the manga ID used to make it unique in the database
	ID string `json:"id"`

	// Title is the title of the manga
	Title string `json:"title"`

	// Author is the author of the manga
	Author string `json:"author"`

	// Magazine is the magazine the manga was published in
	Magazine string `json:"magazine"`

	// Publisher is the publisher of the manga
	Publisher string `json:"publisher"`

	// Volumes is a slice of volumes
	Volumes []Volume `json:"volumes"`
}

// Volume is a struct that represents a volume of a manga
type Volume struct {
	// ID is the volume ID used to make it unique in the database
	ID string `json:"id"`

	// MangaID is the manga ID that the volume belongs to
	MangaID string `json:"mangaID" db:"mangaID"`

	// Number is the volume number
	Number int `json:"number"`

	// Title is the title of the volume
	Title string `json:"title"`

	// ReleaseDate is the date the volume was released
	ReleaseDate string `json:"releaseDate" db:"releaseDate"`

	// ISBN is the International Standard Book Number
	ISBN string `json:"isbn"`

	// Chapters is a slice of chapters
	Chapters []Chapter `json:"chapters"`
}

// Chapter is a struct that represents a chapter of a manga
type Chapter struct {
	// ID is the chapter ID used to make it unique in the database
	ID string `json:"id"`

	// VolumeID is the volume ID that the chapter belongs to
	VolumeID string `json:"volumeID"`

	// Number is the chapter number
	Number int `json:"number"`

	// Title is the title of the chapter
	Title string `json:"title"`
}

// MangaStore is a struct that represents the manga store
type MangaStore struct {
	DataBase *sqlx.DB
}

// NewMangaStore is a function that returns a new manga store
func NewMangaStore(db *sqlx.DB) *MangaStore {
	return &MangaStore{db}
}

// GetManga is a function that returns a manga by ID
func (s *MangaStore) GetManga(id string) (Manga, error) {
	var manga Manga
	err := s.DataBase.Get(&manga, "SELECT * FROM manga WHERE id = $1", id)
	return manga, err
}

// ListManga is a function that returns a list of all manga
func (s *MangaStore) ListManga() ([]Manga, error) {
	var manga []Manga
	err := s.DataBase.Select(&manga, "SELECT * FROM manga")
	return manga, err
}

// CreateManga is a function that creates a manga with a generated ID
func (s *MangaStore) CreateManga(manga *Manga) (*Manga, error) {
	manga.ID = utils.NewUUID()

	_, err := s.DataBase.Exec(
		"INSERT INTO manga (id, title, author, magazine, publisher) VALUES ($1, $2, $3, $4, $5)",
		manga.ID, manga.Title, manga.Author, manga.Magazine, manga.Publisher,
	)
	return manga, err
}

// GetVolume is a function that returns a volume of a manga by ID
func (s *MangaStore) GetVolume(id string) (Volume, error) {
	var volume Volume
	err := s.DataBase.Get(&volume, "SELECT * FROM mangaVolumes WHERE id = $1", id)
	return volume, err
}

// ListVolumes is a function that returns a list of all volumes of a manga
func (s *MangaStore) ListVolumes() ([]Volume, error) {
	var volumes []Volume
	err := s.DataBase.Select(&volumes, "SELECT * FROM mangaVolumes")
	return volumes, err
}

// CreateVolume is a function that creates a volume of a manga with a generated ID
func (s *MangaStore) CreateVolume(volume *Volume) (*Volume, error) {
	volume.ID = utils.NewUUID()

	_, err := s.DataBase.Exec(
		"INSERT INTO mangaVolumes (id, mangaID, number, title, releaseDate, isbn) VALUES ($1, $2, $3, $4, $5, $6)",
		volume.ID, volume.MangaID, volume.Number, volume.Title, volume.ReleaseDate, volume.ISBN,
	)
	return volume, err
}
