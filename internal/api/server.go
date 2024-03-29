package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"pasetoservice/internal/appconfig"
	"pasetoservice/internal/auth"
	"pasetoservice/internal/models"
)

type App struct {
	token     *auth.PasetoAuth
	routerApi *fiber.App
	config    *appconfig.Config
}

func NewApp(c *appconfig.Config, routerApi *fiber.App) (*App, error) {

	pasetoToken, err := auth.NewPaseto([]byte(c.TokenKey))
	if err != nil {
		return nil, err
	}

	app := &App{
		token:     pasetoToken,
		routerApi: routerApi,
		config:    c,
	}

	app.SetApi()

	return app, nil
}

func (a *App) SetApi() {
	a.routerApi.Post("/login", a.Login)

	protectedApi := a.routerApi.Group("/api", a.CheckAuth())

	protectedApi.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString(`<h1>Hello</h1>`)
	})

	protectedApi.Get("/account", func(c *fiber.Ctx) error {

		val := c.Locals("claims")

		v, ok := val.(*models.ServiceClaims)

		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "type conversion error"})
		}

		fmt.Printf("%#v", v)

		owner := fmt.Sprintf("<h3>Account owner - %s</h3>", v.Name)
		role := fmt.Sprintf("<h3>Account role - %s</h3>", v.Role)
		footer := fmt.Sprintf("<h3>Account footer - %s</h3>", v.MetaData)

		return c.SendString(owner + role + footer)
	})

}

func (a *App) Start() error {
	return a.routerApi.Listen(a.config.Address)
}
