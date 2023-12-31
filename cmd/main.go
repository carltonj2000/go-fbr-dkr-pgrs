package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"github.com/carltonj2000/go-fbr-dkr-pgrs/database"
)

func main() {
	database.ConnectDb()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine, ViewsLayout: "layouts/main"})
	setupRoutes(app)
	app.Static("/", "./public")
	// app.Use(logger.New(), handlers.NotFound)
	app.Use(logger.New())
	app.Listen(":3000")
}
