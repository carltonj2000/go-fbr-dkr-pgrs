package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/carltonj2000/go-fbr-dkr-pgrs/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Post("/", handlers.CreateFacts)
}
