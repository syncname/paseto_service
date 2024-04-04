package api

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

const (
	authHeader = "Authorization"
	typeBearer = "bearer"
)

var (
	ErrMissingAuthHeader = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("invalid authorization header")
)

func (a *App) CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authVal := c.Get(authHeader)

		if len(authVal) == 0 {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": ErrMissingAuthHeader.Error()})
		}

		splitHeader := strings.Fields(authVal)

		if len(splitHeader) < 2 {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": ErrInvalidAuthHeader.Error()})
		}

		authType := strings.ToLower(splitHeader[0])
		if authType != typeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authType)
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": err.Error()})
		}

		claims, err := a.token.VerifyToken(splitHeader[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": err.Error()})
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}
