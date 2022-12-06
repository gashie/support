package settings

import (
	"strconv"
	"support-ticket/database"
	"support-ticket/models"

	"github.com/gofiber/fiber/v2"
)

func CreateStatus(c *fiber.Ctx) error {
	var model models.Status

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
func AllStatus(c *fiber.Ctx) error {
	var model []models.Status
	database.DB.Find(&model)
  return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    model,
		"Message": "Record Found",
	})
}

func GetStatus(c *fiber.Ctx) error {
	id,_ := strconv.Atoi(c.Params("id"))
	var status models.Status

	database.DB.Find(&status, "id = ?", id)
	if status.Title == "" {
		return c.Status(200).JSON(fiber.Map{
			"Status":  0,
			"Message": "No Record Found",
		})
	}

    return c.Status(200).JSON(fiber.Map{
		"Status":  1,
		"Data":    status,
		"Message": "Record Found",
	})
}
func UpdateStatus(c *fiber.Ctx) error {
	var body models.Status

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  0,
			"Message": "Fields cannot be empty",
		})
	}

	var status models.Status

	Id := c.Params("id")
	// CHECK AVAILABLE STATUS
	err := database.DB.First(&status, "id = ?", Id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Record not found",
			"Status":0,
		})
	}
   // UPDATE USER DATA
	if body.Title != "" {
		status.Title = body.Title
	}
	// status.Title = body.Title
	// status.Description = body.Description
	
	errUpdate := database.DB.Model(&status).Updates(body).Error //save new changes

	// errUpdate := database.DB.Save(&status).Error //save new changes
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"Message": "Failed to update record",
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Record Updated",
		"Data":    status,
	})
}

func DeleteStatus(c *fiber.Ctx) error {
	Id := c.Params("id")
	var model models.Status

	//	CHECK AVAILABLE USER
	err := database.DB.Debug().First(&model, "id=?", Id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"Message": "Record not found",
			"Status":0,
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
		"Status":    1,
	})
}