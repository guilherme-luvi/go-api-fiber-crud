package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/handlers"
)

func initRoutes(router *fiber.App) {

	// Initialize handlers
	handlers.InitHandler()

	v1 := router.Group("/api/v1")
	{
		// User routes
		v1.Get("/users", handlers.GetUsers)
	}

}