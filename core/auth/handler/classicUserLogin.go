package authHandler

import (
	"strings"

	sharedDomain "api.ruxwez.dev/core/_shared"
	authApplication "api.ruxwez.dev/core/auth/application"
	authDomain "api.ruxwez.dev/core/auth/domain"
	emailLib "api.ruxwez.dev/libs/email"
	"github.com/gofiber/fiber/v2"
)

func ClassicUserLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body authDomain.ClassicLoginBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: sharedDomain.ErrInvalidRequestBody.Error(),
			})
		}

		// Limpiamos los campos del body
		body.Email = strings.TrimSpace(body.Email)
		body.Password = strings.TrimSpace(body.Password)

		// Validamos los campos
		if body.Email == "" {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: authDomain.ErrEmailEmpty.Error(),
			})
		}

		if !emailLib.CheckFormat(body.Email) {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: authDomain.ErrEmailInvalid.Error(),
			})
		}

		if body.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: authDomain.ErrPasswordEmpty.Error(),
			})
		}

		token, err := authApplication.ClassicUserLogin(body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: err.Error(),
			})
		}

		return c.JSON(sharedDomain.Response{
			Success: true,
			Message: "User created",
			Data:    fiber.Map{"token": token},
		})

	}
}
