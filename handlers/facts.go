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

func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Fact added successfully",
		"Subtitle": "Add more wonderful facts to",
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)
	return ConfirmationView(c)
}
