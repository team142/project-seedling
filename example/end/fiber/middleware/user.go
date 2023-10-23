package middleware

import (
	"github.com/gofiber/fiber/v2"
	end "github.com/team142/project-seedling/example/end/fiber"
	"github.com/team142/project-seedling/example/end/fiber/presenter/v1"
)

// VerifyUserBody will verify if the body passed is valid, setting "User" in fiber Locals to the user
// This should only be used where a BODY with a single `User` is required
func VerifyUserBody() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := &end.User{}
		err := c.BodyParser(user)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(
				presenter.DataError{
					Message: "bad body",
					Data:    string(c.Body()),
					Error:   err,
				},
			)
		}
		c.Locals("User", user)
		return c.Next()
	}
}
