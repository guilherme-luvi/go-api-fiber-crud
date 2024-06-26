package middlewares

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v3"
)

func RequireAuth(c fiber.Ctx) error {
	// Get the token from the request header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "No token provided in the Authorization header",
		})
		return nil
	}

	// Validate Bearer format
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token format",
		})
		return nil
	}

	// Decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the alg:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
			return nil, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token expired",
			})
			return nil
		}

		// Get the user ID
		userID := claims["userId"].(float64)
		userIDStr := strconv.FormatFloat(userID, 'f', -1, 64)

		// Set the user ID in the context
		c.Locals("userId", userIDStr)

		//Continue
		return c.Next()
	} else {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
		return nil
	}

}
