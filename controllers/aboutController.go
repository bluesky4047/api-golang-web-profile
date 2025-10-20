package controllers

import (
	"my-fiber-app/data"
	"my-fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetAbout(c *fiber.Ctx) error {
	response := models.Response{
		Success: true,
		Message: "Data about fetched successfully",
		Data:    data.GetAbout(),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}