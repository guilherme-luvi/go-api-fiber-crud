package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/handlers"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/middlewares"
)

const (
	ApiV1BasePath = "/api/v1"
	UsersRoute    = "/users"
	StarWarsRoute = "/starwars"
)

func initRoutes(router *fiber.App) {

	handlers.InitHandlers()

	// API v1 routes group
	v1 := router.Group(ApiV1BasePath)

	// User routes
	userRoutes := v1.Group(UsersRoute)
	userRoutes.Use([]string{"/delete", "/update"}, middlewares.RequireAuth)
	{
		userRoutes.Post("/login", handlers.Login)
		userRoutes.Post("/create", handlers.CreateUser)
		userRoutes.Get("/get", handlers.GetAllUsers)
		userRoutes.Get("/get/:id", handlers.GetUserByID)
		userRoutes.Put("/update/:id", handlers.UpdateUser)
		userRoutes.Delete("/delete/:id", handlers.DeleteUser)
	}

	// Star Wars API integration routes group
	starWarsRoutes := v1.Group(StarWarsRoute)
	{
		starWarsRoutes.Get("/random/people", handlers.GetStarWarsPeople)
		starWarsRoutes.Get("/random/planet", handlers.GetStarWarsPlanet)
	}
}
