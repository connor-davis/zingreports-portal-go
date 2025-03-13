package users

import (
	"github.com/gofiber/fiber/v2"
)

type ViewAllQuery struct {
	Limit  *int `query:"limit"`
	Offset *int `query:"offset"`
}

// ViewAll         godoc
// @Summary     View All
// @Description View all users.
// @Tags        Users
// @Accept      json
// @Produce     json
// @Param       limit  query int false "Limit"
// @Param       offset query int false "Offset"
// @Success     200 {array} postgres.User
// @Failure     400 {string} string "Invalid query parameters."
// @Failure     401 {string} string "Unauthorized."
// @Failure     500 {string} string "Internal Server Error."
// @Router      /users [get]
func (u *UsersRouter) ViewAll(c *fiber.Ctx) error {
	var query ViewAllQuery

	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid query parameters.")
	}

	users, err := u.userService.FindUsers(query.Limit, query.Offset)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// View         godoc
// @Summary     View
// @Description View a user by id.
// @Tags        Users
// @Accept      json
// @Produce     json
// @Param       id path string true "User ID"
// @Success     200 {object} postgres.User
// @Failure     401 {string} string "Unauthorized"
// @Failure     404 {string} string "The user was not found."
// @Failure     500 {string} string "Internal Server Error."
// @Router      /users/{id} [get]
func (u *UsersRouter) View(c *fiber.Ctx) error {
	userId := c.Params("id")

	user, err := u.userService.FindUserById(userId)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if user == nil || user.Id == "" {
		return c.Status(fiber.StatusNotFound).
			SendString("The user was not found.")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
