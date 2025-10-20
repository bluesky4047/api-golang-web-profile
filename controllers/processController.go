package controllers

import (
	"my-fiber-app/data"
	"my-fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetProcess(c *fiber.Ctx) error {
	response := models.Response{
		Success: true,
		Message: "Data process fetched successfully",
		Data:    data.GetProcess(),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}