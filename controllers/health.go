package controllers

import "github.com/gofiber/fiber/v2"

func Health(c *fiber.Ctx)error {
	return c.Status(200).JSON(fiber.Map{
		"Status":1,
		"Message":"Im alive",
	})
}