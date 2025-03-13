package http

import (
	"github.com/connor-davis/zingreports-portal-go/cmd/api/http/authentication"
	"github.com/connor-davis/zingreports-portal-go/cmd/api/http/middleware"
	"github.com/connor-davis/zingreports-portal-go/internal/services"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type HttpRouter struct {
	middleware     *middleware.Middleware
	authentication *authentication.AuthenticationRouter
}

func NewHttpRouter(
	storage *storage.Storage,
	userService *services.UserService,
	poiService *services.PoiService,
) *HttpRouter {
	middleware := middleware.New(storage, userService)

	authentication := authentication.NewAuthenticationRouter(
		storage,
		userService,
	)

	return &HttpRouter{
		middleware:     middleware,
		authentication: authentication,
	}
}

func (h *HttpRouter) LoadRoutes(router fiber.Router) {
	// Authentication Group
	authentication := router.Group("/authentication")

	authentication.Get(
		"/check",
		h.middleware.Authorized(),
		h.authentication.Check,
	)
	authentication.Post(
		"/login",
		h.authentication.Login,
	)

	// MFA Group
	mfa := authentication.Group("/mfa")

	mfa.Get(
		"/enable",
		h.middleware.Authorized(),
		h.authentication.Enable,
	)
	mfa.Post(
		"/verify",
		h.middleware.Authorized(),
		h.authentication.Verify,
	)
	mfa.Patch(
		"/disable",
		h.middleware.Authorized(),
		h.authentication.Disable,
	)
}
