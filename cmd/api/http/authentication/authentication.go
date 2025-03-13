package authentication

import (
	"github.com/connor-davis/zingreports-portal-go/internal/services"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

type AuthenticationRouter struct {
	storage     *storage.Storage
	userService *services.UserService
}

func New(
	storage *storage.Storage,
	userService *services.UserService,
) *AuthenticationRouter {
	return &AuthenticationRouter{
		storage:     storage,
		userService: userService,
	}
}
