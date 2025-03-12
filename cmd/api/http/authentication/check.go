package authentication

import (
	"github.com/connor-davis/zingreports-portal-go/internal/models/postgres"
	"github.com/gofiber/fiber/v2"
)

// Check godoc
// @Summary Check
// @Description Check if user is authenticated
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} postgres.User "User object"
// @Router /authentication/check [get]
func (a *AuthenticationRouter) Check(c *fiber.Ctx) error {
	user := c.Locals("user").(*postgres.User)

	return c.Status(fiber.StatusOK).JSON(&user)
}
