package controllers

import (
	"my-fiber-app/data"
	"my-fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetServices(c *fiber.Ctx) error {
	response := models.Response{
		Success: true,
		Message: "Data services fetched successfully",
		Data:    data.GetServices(),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}