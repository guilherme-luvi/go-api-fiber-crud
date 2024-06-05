package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/repositories"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/schemas"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/security"
)

func Login(c fiber.Ctx) error {
	request := LoginRequest{}
	if err := c.Bind().Body(&request); err != nil {
		logger.Error("Error binding request body")
		sendError(c, fiber.StatusBadRequest, "Invalid request body")
		return nil
	}

	if err := request.validate(); err != nil {
		logger.Error("Invalid request body", err)
		sendError(c, fiber.StatusBadRequest, err.Error())
		return nil
	}

	user := schemas.User{}

	if err := repositories.NewUserRepository(db).GetUserByEmail(request.Email, &user); err != nil {
		logger.Error("Failed to get user", err)
		sendError(c, fiber.StatusNotFound, "User not found")
		return nil
	}

	if err := security.ComparePassword(user.Password, request.Password); err != nil {
		logger.Error("Invalid password", err)
		sendError(c, fiber.StatusUnauthorized, "Invalid password")
		return nil
	}

	token, err := security.GenerateToken(user.ID)
	if err != nil {
		logger.Error("Failed to generate token", err)
		sendError(c, fiber.StatusInternalServerError, "Failed to generate token")
		return nil
	}

	sendSuccess(c, fiber.StatusOK, token)
	return nil
}
