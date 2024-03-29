package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"pasetoservice/internal/api"
	"pasetoservice/internal/appconfig"
)

func main() {

	c, err := appconfig.Load(".env")
	if err != nil {
		log.Fatal("can't load config:", err)
	}

	routerApi := fiber.New()
	app, err := api.NewApp(c, routerApi)
	if err != nil {
		log.Fatal("create app error:", err)
	}
	app.SetApi()
	log.Fatal(app.Start())

}
