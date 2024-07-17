package routes

import (
	"app/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/", controllers.GetApiRoot)
	route.Get("/vmlist", controllers.GetVMList)
	route.Get("/json", controllers.GetJSON)

	webRoute := a.Group("/web")

	webRoute.Get("/", controllers.GetwebRoot)
}
