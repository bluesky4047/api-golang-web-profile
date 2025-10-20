package controllers

import (
	"my-fiber-app/data"
	"my-fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetBlog(c *fiber.Ctx) error {
	response := models.Response{
		Success: true,
		Message: "Data blog fetched successfully",
		Data:    data.GetBlog(),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
