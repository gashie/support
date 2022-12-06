package routes

import (
	"support-ticket/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/health", controllers.Health)
	app.Post("/api/v1/createticket", controllers.CreateTicket)
	
	
	//status routes
	StatusRoute(app)
    
	//category routes
	CategoryRoute(app)

	//priority routes
	PriorityRoute(app)


}
