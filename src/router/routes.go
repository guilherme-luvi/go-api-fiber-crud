package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/handlers"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/middlewares"
)

func initRoutes(router *fiber.App) {

	// Initialize handlers
	handlers.InitHandler()

	v1 := router.Group("/api/v1")
	{
		// Middleware for routes that require authentication
		v1.Use("/users/:id", middlewares.RequireAuth)

		// User routes
		v1.Post("/users", handlers.CreateUser)
		v1.Get("/users", handlers.GetUsers)
		v1.Delete("/users/:id", handlers.DeleteUser)

		// Login route
		v1.Post("/login", handlers.Login)
	}

}
