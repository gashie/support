package settings

import (
	"strconv"
	"support-ticket/database"
	"support-ticket/models"

	"github.com/gofiber/fiber/v2"
)

func CreatePriority(c *fiber.Ctx) error {
	var model models.Priority
	if err := c.BodyParser(&model); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}

	if model.Title == "" {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Title  cannot be empty",
		})
	}

	if model.Description == "" {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Description  cannot be empty",
		})
	}
	result := database.DB.Create(&model)
	if result.RowsAffected != 1 {

		return c.Status(200).JSON(fiber.Map{
			"Status":  0,
			"Message": "Record Not Saved",
		})

	}
	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Saved",
	})

}
func AllPriority(c *fiber.Ctx) error {
	var model []models.Priority
	database.DB.Find(&model)
	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Found",
	})
}

func GetPriority(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var model models.Priority

	database.DB.Find(&model, "id = ?", id)
	if model.Title == "" {
		return c.Status(200).JSON(fiber.Map{
			"Status":  0,
			"Message": "No Record Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Found",
	})
}
func UpdatePriority(c *fiber.Ctx) error {
	var body models.Priority

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}

	var model models.Priority

	Id := c.Params("id")
	// CHECK AVAILABLE CATEGORY
	err := database.DB.First(&model, "id = ?", Id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Record not found",
			"Status":  0,
		})
	}
	// UPDATE USER DATA
	if body.Title != "" {
		model.Title = body.Title
	}

	errUpdate := database.DB.Model(&model).Updates(body).Error //save new changes

	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Failed to update record",
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Record Updated",
		"Data":    model,
	})
}

func DeletePriority(c *fiber.Ctx) error {
	Id := c.Params("id")
	var model models.Priority

	//	CHECK AVAILABLE USER
	err := database.DB.Debug().First(&model, "id=?", Id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Record not found",
			"Status":  0,
		})
	}

	errDelete := database.DB.Debug().Delete(&model).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Failed to update record",
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Record Removed",
		"Status":  1,
	})
}
