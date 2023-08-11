package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/carltonj2000/go-fbr-dkr-pgrs/database"
)

func main() {
	database.ConnectDb()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine, ViewsLayout: "layouts/main"})
	setupRoutes(app)
	app.Listen(":3000")
}
