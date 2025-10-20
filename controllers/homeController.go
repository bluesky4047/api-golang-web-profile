package controllers

import (
	"my-fiber-app/data"
	"my-fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	response := models.Response{
		Success: true,
		Message: "Data home fetched successfully",
		Data:    data.GetHome(),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
