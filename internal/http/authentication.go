package http

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/connor-davis/zingreports-portal-go/internal/helpers"
	"github.com/gofiber/fiber/v2"
)

type Login struct {
	Email    string `json:"email" validate:"required" binding:"required"`
	Password string `json:"password" validate:"required" binding:"required"`
}

func (h *HttpRouter) Login(c *fiber.Ctx) error {
	var login Login

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid payload.")
	}

	if err := helpers.Validate(login); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid payload.")
	}

	user, err := h.userService.FindUserByEmail(login.Email)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).
			SendString("Invalid email or password.")
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				SendString("Invalid email or password.")
		}

		user.MfaVerified = false

		h.storage.Postgres.
			Updates(&user)

		return c.SendStatus(fiber.StatusOK)
	}
}
