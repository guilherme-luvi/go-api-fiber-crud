package router

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/config"
)

func SetupRouter() {
	app := fiber.New()

	initRoutes(app)

	log.Fatal(app.Listen(":" + config.Port))
}
