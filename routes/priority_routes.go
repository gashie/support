package routes

import (
	"support-ticket/controllers/settings"

	"github.com/gofiber/fiber/v2"
)

func PriorityRoute(app *fiber.App) {

	//priority routes
	app.Post("/api/v1/priority", settings.CreatePriority)
	app.Get("/api/v1/priority/:id", settings.GetPriority)
	app.Get("/api/v1/priority", settings.AllPriority)
	app.Put("/api/v1/priority/:id", settings.UpdatePriority)
	app.Delete("/api/v1/priority/:id", settings.DeletePriority)
}
