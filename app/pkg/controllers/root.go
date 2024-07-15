package controllers

import (
	"github.com/gofiber/fiber/v3"
)

// GetApiRoot func for api/v1/
func GetApiRoot(c fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

// GetJSON func for api/v1/json/
func GetJSON(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":      "Json ENDPOINT",
		"some_key": "some data val",
	})
}
