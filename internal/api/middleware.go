package api

import (
	"github.com/gofiber/fiber/v2"
)

const (
	authHeader = "Authorization"
)

func (a *App) CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authVal := c.Get(authHeader)
		claims, err := a.token.VerifyToken(authVal)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}
