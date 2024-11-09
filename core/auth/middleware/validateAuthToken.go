package authMiddleware

import (
	"strings"

	sharedDomain "internal/core/_shared"
	authInfrastructure "internal/core/auth/infrastructure"

	"github.com/gofiber/fiber/v2"
)

// Middleware para validar el token de autenticación
func ValidateAuthToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(sharedDomain.Response{
			Success: false,
			Message: sharedDomain.ErrUnauthoried.Error(),
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(sharedDomain.Response{
			Success: false,
			Message: sharedDomain.ErrUnauthoried.Error(),
		})
	}

	userId := authInfrastructure.CheckToken(token)

	if userId == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(sharedDomain.Response{
			Success: false,
			Message: sharedDomain.ErrUnauthoried.Error(),
		})
	}

	c.Locals("user_id", userId)

	return c.Next()
}
