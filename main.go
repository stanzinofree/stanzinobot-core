package main

import (
	"botcore/initializer"
	"botcore/mock"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Variable that contain the config values from the .env file trhough initializer modules
var configuration initializer.Config

// init is a function that initializes the application's configuration and connects to the database.
// It loads the configuration from a file in the current directory, and uses it to connect to the database.
// The configuration is also stored in a global variable for later use.
func init() {
	// Load configuration from file.
	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}

	// Connect to database using configuration.
	initializer.ConnectDB(&config)

	// Store configuration in global variable for later use.
	configuration = config
}

// main function initializes the server and sets up the endpoint routes
func main() {
	// Initialize Fiber instance
	core := fiber.New()

	// Add logger and CORS middleware
	core.Use(logger.New())
	core.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	// Healthchecker endpoint
	core.Get("/healthchecker", func(c *fiber.Ctx) error {
		// Set initial status for broker and update server - this check will be query other microservices in future
		broker := "UP"
		updateServer := "UP"
		// Return status 200 and JSON response
		return c.Status(200).JSON(fiber.Map{
			"status":       "success",
			"broker":       broker,
			"updateServer": updateServer,
		})
	})

	// Root endpoint
	core.Get("/", func(c *fiber.Ctx) error {
		// Initialize variables for configuration values
		update := ""
		name := configuration.Name
		codename := configuration.Codename
		main := configuration.Main
		version := configuration.Version
		release := configuration.Release
		channel := configuration.Channel
		// Format thisversion and lastVersion with main, version, and release values
		thisversion := fmt.Sprintf("%d.%d.%d", main, version, release)
		retrieveLastVersion := mock.CheckUpdate()
		lastVersion := fmt.Sprintf("%d.%d.%d", retrieveLastVersion.Data.Main, retrieveLastVersion.Data.Version, retrieveLastVersion.Data.Release)
		// Check if an update is available and set update variable accordingly
		if main < retrieveLastVersion.Data.Main {
			update = fmt.Sprintf("Main update available: %s", lastVersion)
		} else if version < retrieveLastVersion.Data.Version {
			update = fmt.Sprintf("Version update available: %s", lastVersion)
		} else if release < retrieveLastVersion.Data.Release {
			update = fmt.Sprintf("Release update available: %s", lastVersion)
		} else {
			update = "No update available"
		}
		// Return status 200 and JSON response
		return c.Status(200).JSON(fiber.Map{
			"status": "success",
			"data": map[string]string{
				"name":     name,
				"codename": codename,
				"version":  thisversion,
				"channel":  channel,
				"update":   update,
			},
		})
	})

	// Start server on specified port
	listner := fmt.Sprintf(":%d", configuration.Port)
	log.Fatal(core.Listen(listner))
}
