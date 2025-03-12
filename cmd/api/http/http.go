package http

import (
	"github.com/connor-davis/zingreports-portal-go/cmd/api/http/authentication"
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
	a := authentication.NewAuthenticationRouter(
		h.storage,
		h.userService,
		h.poiService,
	)
	authentication := router.Group("/authentication")

	authentication.Get(
		"/check",
		a.Authorized(),
		a.Check,
	)

	authentication.Post(
		"/login",
		a.Login,
	)
}
