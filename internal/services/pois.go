package services

import (
	"fmt"

	"github.com/connor-davis/zingreports-portal-go/internal/models/postgres"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

func NewPoiService(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreatePoi(poi postgres.Poi) error {
	result := s.storage.P.Create(&poi)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) UpdatePoiById(id string, poi postgres.Poi) error {
	var old postgres.Poi

	result := s.storage.P.Where("id = $1").Find(&old)

	if result.Error != nil {
		return result.Error
	}

	if poi.Id == "" {
		return fmt.Errorf("The poi was not found.")
	}

	result = s.storage.P.Model(&old).Updates(&poi)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) DeletePoiById(id string) error {
	var poi postgres.Poi

	result := s.storage.P.Where("id = $1").Find(&poi)

	if result.Error != nil {
		return result.Error
	}

	if poi.Id == "" {
		return fmt.Errorf("The poi was not found.")
	}

	result = s.storage.P.Delete(&poi)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) FindPoiById(id string) (*postgres.Poi, error) {
	var poi postgres.Poi

	result := s.storage.P.Where("id = $1", id).Find(&poi)

	if result.Error != nil {
		return nil, result.Error
	}

	if poi.Id == "" {
		return nil, fmt.Errorf("The poi was not found.")
	}

	return &poi, nil
}

func (s *Service) FindPois() ([]postgres.Poi, error) {
	var pois []postgres.Poi

	result := s.storage.P.Find(&pois)

	if result.Error != nil {
		return nil, result.Error
	}

	return pois, nil
}
