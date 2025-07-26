package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Middlewares(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authToken := c.Get("Authorization")
		if authToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization Header"})
		}
		// Remove "Bearer " prefix
		if strings.HasPrefix(authToken, "Bearer ") {
			authToken = strings.TrimPrefix(authToken, "Bearer ")
		}
		dataDecode, err := Decoder(authToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}
		if dataDecode.Role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "This User not Permitted for this function"})
		}

		return c.Next()
	}
}
