package routes

import (
	"api.ruxwez.dev/core/auth"
	"github.com/gofiber/fiber/v2"
)

func apiInit(app *fiber.App) {
	apiGroup := app.Group("/api")
	{
		auth.Routes(apiGroup)
	}
}
