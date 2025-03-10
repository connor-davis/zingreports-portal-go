package services

import (
	"fmt"

	"github.com/connor-davis/zingreports-portal-go/internal/models/postgres"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

func NewUserService(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateUser(user postgres.User) error {
	result := s.storage.P.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) UpdateUserById(id string, user postgres.User) error {
	var old postgres.User

	result := s.storage.P.Where("id = $1").Find(&old)

	if result.Error != nil {
		return result.Error
	}

	if old.Id == "" {
		return fmt.Errorf("The user was not found.")
	}

	result = s.storage.P.Model(&old).Updates(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) DeleteUserById(id string) error {
	var old postgres.User

	result := s.storage.P.Where("id = $1").Find(&old)

	if result.Error != nil {
		return result.Error
	}

	if old.Id == "" {
		return fmt.Errorf("The user was not found.")
	}

	result = s.storage.P.Delete(&old)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) FindUserById(id string) (*postgres.User, error) {
	var user postgres.User

	result := s.storage.P.Where("id = $1", id).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if user.Id == "" {
		return nil, fmt.Errorf("The user was not found.")
	}

	return &user, nil
}

func (s *Service) FindUsers() ([]postgres.User, error) {
	var users []postgres.User

	result := s.storage.P.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
