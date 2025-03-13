package middleware

import (
	"github.com/connor-davis/zingreports-portal-go/internal/services"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

type Middleware struct {
	storage     *storage.Storage
	userService *services.UserService
}

func New(storage *storage.Storage, userService *services.UserService) *Middleware {
	return &Middleware{
		storage:     storage,
		userService: userService,
	}
}
