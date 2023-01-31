package services

import "github.com/stianfro/wotapi/models"

// Service is an interface for the MangaService
type Service struct {
	MangaStore *models.MangaStore
}

// NewService creates a new service with the provided MangaStore
func NewService(ms *models.MangaStore) *Service {
	return &Service{ms}
}

// GetManga implements the GetManga method of the MangaStore interface
func (s *Service) GetManga(mangaID string) (models.Manga, error) {
	return s.MangaStore.GetManga(mangaID)
}

// ListManga implements the ListManga method of the MangaStore interface
func (s *Service) ListManga() ([]models.Manga, error) {
	return s.MangaStore.ListManga()
}

// CreateManga implements the CreateManga method of the MangaStore interface
func (s *Service) CreateManga(manga *models.Manga) (*models.Manga, error) {
	return s.MangaStore.CreateManga(manga)
}

// GetVolume implements the GetVolume method of the MangaStore interface
func (s *Service) GetVolume(volumeID string) (models.Volume, error) {
	return s.MangaStore.GetVolume(volumeID)
}

// ListVolumes implements the ListVolumes method of the MangaStore interface
func (s *Service) ListVolumes(mangaID string) ([]models.Volume, error) {
	return s.MangaStore.ListVolumes(mangaID)
}

// CreateVolume implements the CreateVolume method of the MangaStore interface
func (s *Service) CreateVolume(volume *models.Volume) (*models.Volume, error) {
	return s.MangaStore.CreateVolume(volume)
}
