// cmd/routes.go

package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/diegocheca/goPosts.git/handlers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/", handlers.ListFacts)
    app.Get("/fact", handlers.NewFactView) // Add new route for new view
    app.Post("/fact", handlers.CreateFact)
    app.Get("/fact/:id", handlers.ShowFact)
    // Display `Edit` form
    app.Get("/fact/:id/edit", handlers.EditFact)
    // Update fact
    app.Patch("/fact/:id", handlers.UpdateFact)

    // Delete fact
    app.Delete("/fact/:id", handlers.DeleteFact)
    
}