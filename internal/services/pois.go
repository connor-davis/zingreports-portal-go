package services

import (
	"fmt"

	"github.com/connor-davis/zingreports-portal-go/internal/models/postgres"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

type PoiService struct {
	storage *storage.Storage
}

func NewPoiService(storage *storage.Storage) *PoiService {
	return &PoiService{
		storage: storage,
	}
}

func (s *PoiService) CreatePoi(poi postgres.Poi) error {
	result := s.storage.Postgres.
		Create(&poi)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *PoiService) UpdatePoiById(id string, poi postgres.Poi) error {
	var old postgres.Poi

	result := s.storage.Postgres.
		Where("id = $1").
		Find(&old)

	if result.Error != nil {
		return result.Error
	}

	if poi.Id == "" {
		return fmt.Errorf("The poi was not found.")
	}

	result = s.storage.Postgres.
		Model(&old).
		Updates(&poi)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *PoiService) DeletePoiById(id string) error {
	var poi postgres.Poi

	result := s.storage.Postgres.
		Where("id = $1").
		Find(&poi)

	if result.Error != nil {
		return result.Error
	}

	if poi.Id == "" {
		return fmt.Errorf("The poi was not found.")
	}

	result = s.storage.Postgres.
		Delete(&poi)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *PoiService) FindPoiById(id string) (*postgres.Poi, error) {
	var poi postgres.Poi

	result := s.storage.Postgres.
		Where("id = $1", id).
		Find(&poi)

	if result.Error != nil {
		return nil, result.Error
	}

	if poi.Id == "" {
		return nil, fmt.Errorf("The poi was not found.")
	}

	return &poi, nil
}

func (s *PoiService) FindPois() ([]postgres.Poi, error) {
	var pois []postgres.Poi

	result := s.storage.Postgres.
		Find(&pois)

	if result.Error != nil {
		return nil, result.Error
	}

	return pois, nil
}
