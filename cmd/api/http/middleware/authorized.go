package middleware

import (
	"strings"

	"github.com/connor-davis/zingreports-portal-go/internal/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (m *Middleware) Authorized() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := m.storage.Sessions.Get(c)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				SendString(constants.UnauthorizedMessage)
		}

		userId := session.Get("user").(string)

		if userId == "" {
			return c.Status(fiber.StatusUnauthorized).
				SendString(constants.UnauthorizedMessage)
		}

		log.Infof("🔐 Authorized User with Id: %s", userId)

		err = session.Reset()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error regenerating session.")
		}

		session.Set("user", userId)
		err = session.Save()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error saving session.")
		}

		user, err := m.userService.FindUserById(userId)

		if err != nil && strings.Contains(err.Error(), "The user was not found.") {
			return c.Status(fiber.StatusUnauthorized).
				SendString(constants.UnauthorizedMessage)
		}

		c.Locals("userId", user.Id)
		c.Locals("user", user)

		return c.Next()
	}
}
