package http

import (
	"github.com/connor-davis/zingreports-portal-go/internal/services"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type HttpRouter struct {
	storage     *storage.Storage
	userService *services.UserService
	poiService  *services.PoiService
}

func (h *HttpRouter) LoadRoutes(router fiber.Router) {
	// Authentication Routes
	router.Post("/login", h.Login)
}
