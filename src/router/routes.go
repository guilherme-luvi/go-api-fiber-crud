package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/handlers"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/middlewares"
)

func initRoutes(router *fiber.App) {

	handlers.InitHandlers()

	// Login route
	router.Post("api/login", handlers.Login)

	// API v1 routes group
	v1 := router.Group("/api/v1/")
	{
		// Authentication middleware for routes starting with /users/delete and /users/update
		v1.Use([]string{"users/delete", "users/update"}, middlewares.RequireAuth)

		// User routes
		v1.Post("users/create", handlers.CreateUser)
		v1.Get("users/get", handlers.GetAllUsers)
		v1.Get("users/get/:id", handlers.GetUserByID)
		v1.Put("users/update/:id", handlers.UpdateUser)
		v1.Delete("users/delete/:id", handlers.DeleteUser)

		// Star Wars API integration routes group
		v1.Get("/random/people", handlers.GetStarWarsPeople)
		v1.Get("/random/planet", handlers.GetStarWarsPlanet)
	}
}
