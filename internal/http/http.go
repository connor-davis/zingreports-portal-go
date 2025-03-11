package http

import (
	"github.com/connor-davis/zingreports-portal-go/internal/http/authentication"
	"github.com/connor-davis/zingreports-portal-go/internal/services"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type HttpRouter struct {
	storage     *storage.Storage
	userService *services.UserService
	poiService  *services.PoiService
}

func NewHttpRouter(
	storage *storage.Storage,
	userService *services.UserService,
	poiService *services.PoiService,
) *HttpRouter {
	return &HttpRouter{
		storage:     storage,
		userService: userService,
		poiService:  poiService,
	}
}

func (h *HttpRouter) LoadRoutes(router fiber.Router) {
	// Authentication
	authentication := authentication.NewAuthenticationRouter(
		h.storage,
		h.userService,
		h.poiService,
	)

	router.Post("/login", authentication.Login)
}
