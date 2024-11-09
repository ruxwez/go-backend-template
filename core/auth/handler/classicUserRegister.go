package authHandler

import (
	"strings"

	sharedDomain "api.ruxwez.dev/core/_shared"
	authApplication "api.ruxwez.dev/core/auth/application"
	authDomain "api.ruxwez.dev/core/auth/domain"
	emailLib "api.ruxwez.dev/libs/email"
	passwordLib "api.ruxwez.dev/libs/password"
	usernameLib "api.ruxwez.dev/libs/username"
	"github.com/gofiber/fiber/v2"
)

func ClassicUserRegister() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body authDomain.ClassicRegisterBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: sharedDomain.ErrInvalidRequestBody.Error(),
			})
		}

		// Limpiamos los campos del body
		body.Email = strings.TrimSpace(body.Email)
		body.Password = strings.TrimSpace(body.Password)
		body.Username = strings.TrimSpace(body.Username)
		body.FirstName = strings.TrimSpace(body.FirstName)
		body.LastName = strings.TrimSpace(body.LastName)

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

		if !passwordLib.CheckFormat(body.Password) {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: authDomain.ErrPasswordFormat.Error(),
			})
		}

		if body.Username == "" {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: authDomain.ErrUsernameEmpty.Error(),
			})
		}

		if !usernameLib.CheckFormat(body.Username) {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: authDomain.ErrUsernameFormat.Error(),
			})
		}

		if body.FirstName == "" {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: authDomain.ErrFirstNameEmpty.Error(),
			})
		}

		if body.LastName == "" {
			return c.Status(fiber.StatusBadRequest).JSON(sharedDomain.Response{
				Success: false,
				Message: authDomain.ErrLastNameEmpty.Error(),
			})
		}

		// Ejecutamos el servicio de creacion de usuario

		token, err := authApplication.ClassicUserRegister(body)

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
