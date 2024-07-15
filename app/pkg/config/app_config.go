package config

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

// AppConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func AppConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
