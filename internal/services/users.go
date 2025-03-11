package services

import (
	"fmt"

	"github.com/connor-davis/zingreports-portal-go/internal/models/postgres"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

type UserService struct {
	storage *storage.Storage
}

func NewUserService(storage *storage.Storage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s *UserService) CreateUser(user postgres.User) error {
	result := s.storage.Postgres.
		Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *UserService) UpdateUserById(id string, user postgres.User) error {
	var old postgres.User

	result := s.storage.Postgres.
		Where("id = $1").
		Find(&old)

	if result.Error != nil {
		return result.Error
	}

	if old.Id == "" {
		return fmt.Errorf("The user was not found.")
	}

	result = s.storage.Postgres.
		Model(&old).
		Updates(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *UserService) DeleteUserById(id string) error {
	var old postgres.User

	result := s.storage.Postgres.
		Where("id = $1").
		Find(&old)

	if result.Error != nil {
		return result.Error
	}

	if old.Id == "" {
		return fmt.Errorf("The user was not found.")
	}

	result = s.storage.Postgres.
		Delete(&old)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *UserService) FindUserById(id string) (*postgres.User, error) {
	var user postgres.User

	result := s.storage.Postgres.
		Where("id = $1", id).
		Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if user.Id == "" {
		return nil, fmt.Errorf("The user was not found.")
	}

	return &user, nil
}

func (s *UserService) FindUserByEmail(email string) (*postgres.User, error) {
	var user postgres.User

	result := s.storage.Postgres.
		Where("email = $1", email).
		Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if user.Id == "" {
		return nil, fmt.Errorf("The user was not found.")
	}

	return &user, nil
}

func (s *UserService) FindUsers() ([]postgres.User, error) {
	var users []postgres.User

	result := s.storage.Postgres.
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
