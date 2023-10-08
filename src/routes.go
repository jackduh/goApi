package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackduh/goApi/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/facts", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
}
