package main

import (
	"app/pkg/config"
	"app/pkg/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	_ "github.com/joho/godotenv/autoload"
)

var serveParams = fmt.Sprintf(
	"%s:%s",
	os.Getenv("SERVER_HOST"),
	os.Getenv("SERVER_PORT"),
)

func main() {
	// Define Fiber config.
	appConfig := config.AppConfig()

	// Define a new Fiber app with config.
	app := fiber.New(appConfig)

	// redirect / to api/v1
	app.Get("/", func(c fiber.Ctx) error {
		return c.Redirect().To("/api/v1")
	})

	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(serveParams))
}
