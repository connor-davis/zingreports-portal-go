package authentication

import (
	"bytes"
	"encoding/base32"
	"image/png"
	"log"

	"github.com/connor-davis/zingreports-portal-go/internal/constants"
	"github.com/connor-davis/zingreports-portal-go/internal/models/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type VerifyMfaPayload struct {
	Code string `json:"code" binding:"required"`
}

type DisableMfaQuery struct {
	UserId string `query:"userId" binding:"required"`
}

// Enable       godoc
// @Summary     Enable
// @Description Enable MFA for the current user.
// @Tags        Authentication
// @Accept      json
// @Produce     png
// @Success     200 {file} png "QR Code"
// @Failure     500 {string} string "Internal Server Error"
// @Router      /authentication/mfa/enable [get]
func (a *AuthenticationRouter) Enable(c *fiber.Ctx) error {
	user := c.Locals("user").(*postgres.User)

	log.Printf("Old: %v", user)

	if user.MfaSecret == "" {
		secret, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "Zing Fibre",
			AccountName: user.Email,
			Period:      30,
			Digits:      otp.DigitsSix,
			Algorithm:   otp.AlgorithmSHA1,
			SecretSize:  32,
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString(err.Error())
		}

		secretString := secret.Secret()

		user.MfaSecret = secretString

		a.storage.Postgres.
			Updates(user)
	}

	log.Printf("New: %v", user)

	secretBytes, err := base32.StdEncoding.WithPadding(base32.NoPadding).
		DecodeString(user.MfaSecret)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Zing Fibre",
		AccountName: user.Email,
		Period:      30,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
		Secret:      secretBytes,
		SecretSize:  32,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	var pngBuffer bytes.Buffer

	image, err := secret.Image(256, 256)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	png.Encode(&pngBuffer, image)

	c.Response().Header.Set("Content-Type", "image/png")

	return c.Status(fiber.StatusOK).Send(pngBuffer.Bytes())
}

// Verify       godoc
// @Summary     Verify
// @Description Verify MFA for the current user.
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       payload body VerifyMfaPayload true "Verify MFA Payload."
// @Success     200 {string} string "ok"
// @Failure     400 {string} string "Invalid request body."
// @Failure     400 {string} string "MFA has not been enabled."
// @Failure     400 {string} string "Invalid MFA code."
// @Failure     500 {string} string "Internal Server Error."
// @Router      /authentication/mfa/verify [post]
func (a *AuthenticationRouter) Verify(c *fiber.Ctx) error {
	user := c.Locals("user").(*postgres.User)

	var payload VerifyMfaPayload

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid request body.")
	}

	if user.MfaSecret == "" {
		return c.Status(fiber.StatusBadRequest).
			SendString("MFA has not been enabled.")
	}

	log.Printf("Code: %s", payload.Code)
	log.Printf("User: %v", user)

	if !totp.Validate(payload.Code, user.MfaSecret) {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid MFA code.")
	}

	user.MfaEnabled = true
	user.MfaVerified = true

	a.storage.Postgres.Updates(user)

	return c.Status(fiber.StatusOK).SendString("ok")
}

// Disable      godoc
// @Summary     Disable
// @Description Disable MFA for the current user.
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       userId query string true "User ID"
// @Success     200 {string} string "ok"
// @Failure     400 {string} string "Invalid query parameters."
// @Failure     401 {string} string "Unauthorized."
// @Failure     403 {string} string "Forbidden."
// @Failure     404 {string} string "The user was not found."
// @Router      /authentication/mfa/disable [patch]
func (a *AuthenticationRouter) Disable(c *fiber.Ctx) error {
	var query DisableMfaQuery

	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid query parameters.")
	}

	currentUser := c.Locals("user").(*postgres.User)

	if currentUser.Role != "admin" {
		return c.Status(fiber.StatusForbidden).
			SendString(constants.UnauthorizedMessage)
	}

	if query.UserId == "" {
		return c.Status(fiber.StatusNotFound).
			SendString("The user was not found.")
	}

	var user postgres.User

	a.storage.Postgres.
		Where("id = $1", query.UserId).
		Find(&user)

	if user.Id == "" {
		return c.Status(fiber.StatusNotFound).
			SendString("The user was not found.")
	}

	a.storage.Postgres.
		Model(user).
		Updates(map[string]interface{}{
			"mfa_enabled":  false,
			"mfa_verified": false,
			"mfa_secret":   nil,
		})

	return c.Status(fiber.StatusOK).SendString("ok")
}
