package handlers

import (
	"fmt"

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
	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}
	return c.Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact":  fact,
		"Id":    id,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}

func EditFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")
	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}
	return c.Render("edit", fiber.Map{
		"Title": "Edit Fact",
		"Fact":  fact,
	})
}

func UpdateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")
	fmt.Printf("id %+v\n", id)

	if err := c.BodyParser(fact); err != nil {
		fmt.Printf("error %+v\n", err.Error())
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}
	fmt.Printf("fact %+v\n", fact)

	result := database.DB.Db.Where("id = ?", id).Updates(fact)
	if result.Error != nil {
		return EditFact(c)
	}
	return ShowFact(c)
}

func DeleteFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")
	fmt.Printf("id %+v\n", id)

	result := database.DB.Db.Where("id = ?", id).Delete(&fact)
	if result.Error != nil {
		return NotFound(c)
	}
	return ListFacts(c)
}
