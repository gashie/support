package routes

import (
	"support-ticket/controllers/settings"

	"github.com/gofiber/fiber/v2"
)

func StatusRoute(app *fiber.App) {

	//status routes
	app.Post("/api/v1/status", settings.CreateStatus)
	app.Get("/api/v1/status/:id", settings.GetStatus)
	app.Get("/api/v1/status", settings.AllStatus)
	app.Put("/api/v1/status/:id", settings.UpdateStatus)
	app.Delete("/api/v1/status/:id", settings.DeleteStatus)
}
