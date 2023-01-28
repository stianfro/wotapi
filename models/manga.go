package models

// Manga is a struct that represents a manga
type Manga struct {
	// ID is the manga ID used to make it unique in the database
	ID string `json:"id"`

	// Title is the title of the manga
	Title string `json:"title" binding:"required"`

	// Author is the author of the manga
	Author string `json:"author" binding:"required"`

	// Magazine is the magazine the manga was published in
	Magazine string `json:"magazine" binding:"required"`

	// Publisher is the publisher of the manga
	Publisher string `json:"publisher" binding:"required"`

	// Volumes is a slice of volumes
	Volumes []Volume `json:"volumes"`
}

// Volume is a struct that represents a volume of a manga
type Volume struct {
	// ID is the volume ID used to make it unique in the database
	ID int `json:"id"`

	// MangaID is the manga ID that the volume belongs to
	MangaID string `json:"mangaID"`

	// Number is the volume number
	Number int `json:"number"`

	// Title is the title of the volume
	Title string `json:"title"`

	// ReleaseDate is the date the volume was released
	ReleaseDate string `json:"releaseDate"`

	// ISBN is the International Standard Book Number
	ISBN string `json:"isbn"`

	// Chapters is a slice of chapters
	Chapters []Chapter `json:"chapters"`
}

// Chapter is a struct that represents a chapter of a manga
type Chapter struct {
	// ID is the chapter ID used to make it unique in the database
	ID int `json:"id"`

	// VolumeID is the volume ID that the chapter belongs to
	VolumeID int `json:"volumeID"`

	// Number is the chapter number
	Number int `json:"number"`

	// Title is the title of the chapter
	Title string `json:"title"`
}
