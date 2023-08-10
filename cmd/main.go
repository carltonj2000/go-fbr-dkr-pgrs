package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/carltonj2000/go-fbr-dkr-pgrs/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Carlton Three!")
	})

	app.Listen(":3000")
}
