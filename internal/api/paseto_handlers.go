package api

import (
	"github.com/gofiber/fiber/v2"
	"pasetoservice/internal/models"
)

func (a *App) Login(c *fiber.Ctx) error {
	var creds models.Credentials

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	expectedPass, ok := users[creds.Username]
	if !ok || expectedPass != creds.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "bad password or username"})
	}

	pasetoToken, err := a.token.NewToken(models.TokenData{
		Subject:  "for user",
		Duration: a.config.TokenDuration,
		AdditionalClaims: models.AdditionalClaims{
			Name: creds.Username,
			Role: creds.Username,
		},
		Footer: models.Footer{MetaData: "footer for " + creds.Username},
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": pasetoToken})
}
