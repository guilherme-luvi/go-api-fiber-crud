package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/services"
)

func GetStarWarsPeople(c fiber.Ctx) error {
	resp, err := services.GetRandomStarWarsPeople()
	if err != nil {
		logger.Error("Failed to get Star Wars people", err)
		sendError(c, fiber.StatusInternalServerError, "Failed to get Star Wars people")
		return nil
	}

	sendSuccess(c, fiber.StatusOK, string(resp))
	return nil
}

func GetStarWarsPlanet(c fiber.Ctx) error {
	resp, err := services.GetRandomStarWarsPlanet()
	if err != nil {
		logger.Error("Failed to get Star Wars planet", err)
		sendError(c, fiber.StatusInternalServerError, "Failed to get Star Wars planet")
		return nil
	}

	sendSuccess(c, fiber.StatusOK, string(resp))
	return nil
}
