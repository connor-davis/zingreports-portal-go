package authentication

import (
	"github.com/connor-davis/zingreports-portal-go/internal/services"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

type AuthenticationRouter struct {
	storage     *storage.Storage
	userService *services.UserService
	poiService  *services.PoiService
}

func NewAuthenticationRouter(
	storage *storage.Storage,
	userService *services.UserService,
	poiService *services.PoiService,
) *AuthenticationRouter {
	return &AuthenticationRouter{
		storage:     storage,
		userService: userService,
		poiService:  poiService,
	}
}
