package main

import (
	"app/pkg/config"
	"app/pkg/routes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
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

	logFile := setUpLog(app)
	defer logFile.Close()

	// redirect / to api/v1
	app.Get("/", func(c fiber.Ctx) error {
		return c.Redirect().To("/api/v1")
	})

	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(serveParams))
}

func setUpLog(app *fiber.App) *os.File {
	logFilename := os.Getenv("LOG_FILENAME")

	logFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Create a multi-writer to write to both stdout and the log file
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	app.Use(logger.New(
		logger.Config{
			Output: multiWriter,
		},
	))
	return logFile
}
