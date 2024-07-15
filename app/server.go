package main

import (
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
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		fmt.Println(os.Getenv("SERVER_PORT"))
		return c.SendString("Hello, World 👋!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(serveParams))
}
