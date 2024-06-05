package handlers

import "github.com/gofiber/fiber/v3"

func GetUsers(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status_code": 200,
		"user": fiber.Map{
			"id":    1,
			"name":  "John Doe",
			"email": "123@email.com",
		},
	})
}
