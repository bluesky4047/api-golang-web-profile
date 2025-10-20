package controllers

import (
	"my-fiber-app/data"
	"my-fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetTestimonials(c *fiber.Ctx) error {
	response := models.Response{
		Success: true,
		Message: "Data testimonials fetched successfully",
		Data:    data.GetTestimonials(),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}