package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get To Do List API",
	})
}

func GetByParam(c *fiber.Ctx) error {
	param := c.Params("id")

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Get To Do List API with param %s", param),
	})
}

func Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Create To Do List API",
	})
}

func UpdateByParam(c *fiber.Ctx) error {
	param := c.Params("id")

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Update To Do List API with param %s", param),
	})
}

func BulkDelete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Delete To Do List API",
	})
}
