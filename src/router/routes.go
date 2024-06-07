package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/handlers"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/middlewares"
)

var (
	basePath         = "api/"
	usersBasePath    = basePath + "v1/users"
	starwarsBasePath = basePath + "v1/starwars"
)

func initRoutes(router *fiber.App) {

	handlers.InitHandlers()

	// Login route
	router.Post(basePath+"/login", handlers.Login)

	// User routes group
	v1Users := router.Group(usersBasePath)
	{
		// Authentication middleware for routes starting with /delete and /update
		v1Users.Use([]string{"/delete", "/update"}, middlewares.RequireAuth)

		// User routes
		v1Users.Post("/create", handlers.CreateUser)
		v1Users.Get("/get", handlers.GetAllUsers)
		v1Users.Get("/get/:id", handlers.GetUserByID)
		v1Users.Put("/update/:id", handlers.UpdateUser)
		v1Users.Delete("/delete/:id", handlers.DeleteUser)
	}

	// Star Wars API integration routes group
	v1StarWars := router.Group(starwarsBasePath)
	{
		v1StarWars.Get("/random/people", handlers.GetStarWarsPeople)
		v1StarWars.Get("/random/planet", handlers.GetStarWarsPlanet)
	}
}
