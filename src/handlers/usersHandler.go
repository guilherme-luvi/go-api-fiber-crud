package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/repositories"
	"github.com/guilherme-luvi/go-api-fiber-crud/src/schemas"
)

func CreateUser(c fiber.Ctx) error {
	request := CreateUserRequest{}
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

	user := schemas.User{
		Name:  request.Name,
		Email: request.Email,
	}

	if err := repositories.NewUserRepository(db).CreateUser(user); err != nil {
		logger.Error("Failed to create user", err)
		sendError(c, fiber.StatusInternalServerError, "Failed to create user")
		return nil
	}

	sendSuccess(c, fiber.StatusCreated, user)
	return nil
}

func GetUsers(c fiber.Ctx) error {
	users, err := repositories.NewUserRepository(db).GetUsers()
	if err != nil {
		logger.Error("Failed to get users", err)
		sendError(c, fiber.StatusInternalServerError, "Failed to get users")
		return nil
	}

	sendSuccess(c, fiber.StatusOK, users)
	return nil
}
