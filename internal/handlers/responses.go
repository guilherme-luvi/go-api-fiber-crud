package handlers

import "github.com/gofiber/fiber/v3"

func sendSuccess(c fiber.Ctx, statusCode int, data interface{}) {
	c.Response().Header.Set("Content-Type", "application/json")
	c.Status(statusCode).JSON(fiber.Map{
		"status_code": statusCode,
		"data":        data,
	})
}

func sendError(c fiber.Ctx, statusCode int, message string) {
	c.Response().Header.Set("Content-Type", "application/json")
	c.Status(statusCode).JSON(fiber.Map{
		"status_code": statusCode,
		"message":     message,
	})
}
