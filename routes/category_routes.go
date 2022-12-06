package routes

import (
	"support-ticket/controllers/settings"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoute(app *fiber.App) {

	//status routes
	app.Post("/api/v1/category", settings.CreateCategory)
	app.Get("/api/v1/category/:id", settings.GetCategory)
	app.Get("/api/v1/category", settings.AllCategory)
	app.Put("/api/v1/category/:id", settings.UpdateCategory)
	app.Delete("/api/v1/category/:id", settings.DeleteCategory)
}
