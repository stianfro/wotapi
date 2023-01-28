package models

// Manga is a struct that represents a manga
type Manga struct {
	// ID is the manga ID used to make it unique in the database
	ID int `json:"id"`

	// Title is the title of the manga
	Title string `json:"title"`

	// Author is the author of the manga
	Author string `json:"author"`

	// Magazine is the magazine the manga was published in
	Magazine string `json:"magazine"`

	// Volumes is a slice of volumes
	Volumes []Volume `json:"volumes"`
}

// Volume is a struct that represents a volume of a manga
type Volume struct {
	// ID is the volume ID used to make it unique in the database
	ID int `json:"id"`

	// Number is the volume number
	Number int `json:"number"`

	// Title is the title of the volume
	Title string `json:"title"`

	// ReleaseDate is the date the volume was released
	Release string `json:"releaseDate"`

	// ISBN is the International Standard Book Number
	ISBN string `json:"isbn"`

	// Chapters is a slice of chapters
	Chapters []Chapter `json:"chapters"`
}

// Chapter is a struct that represents a chapter of a manga
type Chapter struct {
	// ID is the chapter ID used to make it unique in the database
	ID int `json:"id"`

	// Number is the chapter number
	Number int `json:"number"`

	// Title is the title of the chapter
	Title string `json:"title"`
}
