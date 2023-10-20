package middleware

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(db *sql.DB, params []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
