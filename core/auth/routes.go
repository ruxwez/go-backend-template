package auth

import (
	authHandler "api.ruxwez.dev/core/auth/handler"
	"github.com/gofiber/fiber/v2"
)

/* Creamos el grupo de auth */
func Routes(api fiber.Router) {
	auth := api.Group("/auth")
	{
		auth.Post("/register", authHandler.ClassicUserRegister()) // Creamos la ruta de registro clasico
		auth.Post("/login", authHandler.ClassicUserLogin())       // Creamos la ruta de login clasico
		auth.Get("/discord", func(c *fiber.Ctx) error {
			return c.SendString("Discord")
		})
	}
}
