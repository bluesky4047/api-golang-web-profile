package controllers

import (
	"my-fiber-app/data"
	"my-fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetContact(c *fiber.Ctx) error {
	response := models.Response{
		Success: true,
		Message: "Data contact fetched successfully",
		Data:    data.GetContact(),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func PostContact(c *fiber.Ctx) error {
	var contact models.ContactRequest

	// Parsing JSON dari body request
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: "Invalid request body",
		})
	}

	// Misal validasi sederhana (pastikan name dan email tidak kosong)
	if contact.Name == "" || contact.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: "Name and email are required",
		})
	}

	// Return response langsung
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Contact data received successfully",
		Data:    contact,
	})
}