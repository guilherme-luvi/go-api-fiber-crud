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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error binding request body",
		})
	}

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := repositories.NewUserRepository(db).CreateUser(user); err != nil {
		logger.Error("Failed to create user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

func GetUsers(c fiber.Ctx) error {
	users, err := repositories.NewUserRepository(db).GetUsers()
	if err != nil {
		logger.Error("Failed to get users", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get users",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}
