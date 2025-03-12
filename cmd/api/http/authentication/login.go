package authentication

import (
	"github.com/connor-davis/zingreports-portal-go/internal/helpers"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email    string `json:"email" validate:"required" binding:"required"`
	Password string `json:"password" validate:"required" binding:"required"`
}

// Login godoc
// @Summary Authenticate User
// @Description Authenticate a user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body Login true "Login payload."
// @Success 200 {string} string "Authenticated."
// @Failure 400 {string} string "Invalid request body."
// @Failure 401 {string} string "Unauthorized."
// @Failure 500 {string} string "Internal Server Error."
// @Router /authentication/login [post]
func (a *AuthenticationRouter) Login(c *fiber.Ctx) error {
	var login Login

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid request body.")
	}

	if err := helpers.Validate(login); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid request body.")
	}

	user, err := a.userService.FindUserByEmail(login.Email)

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

		a.storage.Postgres.
			Updates(&user)

		session, err := a.storage.Sessions.Get(c)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString("Error acquiring session.")
		}

		session.Set("user", user.Id)

		if err := session.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString("Error saving session.")
		}

		return c.Status(fiber.StatusOK).SendString("ok")
	}
}
