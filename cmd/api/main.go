package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ticket-booking-app/handlers"
	"github.com/ticket-booking-app/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "BookYourShow",
		ServerHeader: "Fiber",
	})

	// Initialise Repositories
	eventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api")

	// Initialise Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")

}
