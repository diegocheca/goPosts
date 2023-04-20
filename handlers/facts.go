package handlers

import (

    "github.com/gofiber/fiber/v2"
    "github.com/diegocheca/goPosts.git/models"
    "github.com/diegocheca/goPosts.git/database"
    
)

func ListFacts(c *fiber.Ctx) error {
    facts := []models.Fact{}

    database.DB.Db.Find(&facts)

    //return c.Status(200).JSON(facts)
    return c.Render("index", fiber.Map{
        "Title":  "Div Rhino Trivia Time",
        "Subtitle": "Facts for funtimes with friends!",
        "Facts": facts,
    })
}


// Create new Fact View handler
func NewFactView(c *fiber.Ctx) error {
    return c.Render("new", fiber.Map{
        "Title":    "New Fact",
        "Subtitle": "Add a cool fact!",
    })
}


// 1. New Confirmation view
func ConfirmationView(c *fiber.Ctx) error {
    return c.Render("confirmation", fiber.Map{
        "Title":    "Fact added successfully",
        "Subtitle": "Add more wonderful facts to the list!",
    })
}


func CreateFact(c *fiber.Ctx) error {
    fact := new(models.Fact)
    if err := c.BodyParser(fact); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    database.DB.Db.Create(&fact)

    //return c.Status(200).JSON(fact)
    return ConfirmationView(c)
}