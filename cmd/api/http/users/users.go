package users

import (
	"github.com/connor-davis/zingreports-portal-go/internal/services"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

type UsersRouter struct {
	storage     *storage.Storage
	userService *services.UserService
}

func New(
	storage *storage.Storage,
	userService *services.UserService,
) *UsersRouter {
	return &UsersRouter{
		storage:     storage,
		userService: userService,
	}
}
