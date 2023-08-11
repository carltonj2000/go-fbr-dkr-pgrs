package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/carltonj2000/go-fbr-dkr-pgrs/database"
	"github.com/carltonj2000/go-fbr-dkr-pgrs/models"
)

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Fact",
		"Subtitle": "Add A Fact",
	})
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Render("index", fiber.Map{
		"Title":    "Full Stack Facts App",
		"Subtitle": "Go Lang, Fiber, Postgres, Docker Example",
		"Facts":    facts,
	})
}

func CreateFacts(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return NewFactView(c)
	}

	result := database.DB.Db.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}

	return ListFacts(c)
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")
	database.DB.Db.Where("id = ?", id).First(&fact)
	return c.Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact":  fact,
	})
}
