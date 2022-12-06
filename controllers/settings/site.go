package settings

import (
	"support-ticket/database"
	"support-ticket/models"

	"github.com/gofiber/fiber/v2"
)

func CreateSites(c *fiber.Ctx) error {
	var ticket models.Ticket

	if err := c.BodyParser(&ticket); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}

	if ticket.Subject == "" {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Subject  cannot be empty",
		})
	}

	if ticket.Description == "" {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Description  cannot be empty",
		})
	}
	result := database.DB.Create(&ticket)
	if result.RowsAffected != 1 {

		return c.Status(200).JSON(fiber.Map{
			"Status":  0,
			"Message": "Record Not Saved",
		})

	}
	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    fiber.Map{"description": ticket.Description, "subject": ticket.Subject},
		"Message": "Record Saved",
	})

}
func AllSite(c *fiber.Ctx) error {
	var ticket []models.Ticket
	// database.DB.Preload("Category").Preload("Tasks").Preload("Tasks.Owners").Preload("Tasks.Followers").Preload("WorkFlow").Find(&project)

	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    ticket,
		"Message": "Record Found",
	})
}
