package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func metricsInit(app *fiber.App) {
	app.Get("/metrics", monitor.New())
}
