package controllers

import (
	"app/pkg/virt"
	"os"

	"github.com/gofiber/fiber/v3"
)

var vc = virt.New(os.Getenv("QEMU_HOST"))

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

// GetwebRoot func for api/v1/
func GetwebRoot(c fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

// GetVMList func
func GetVMList(c fiber.Ctx) error {
	vms, _ := vc.ListAllVM()
	return c.Status(fiber.StatusOK).JSON(vms)
}
