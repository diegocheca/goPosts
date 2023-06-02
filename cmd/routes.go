// cmd/routes.go

package main

import (
	"github.com/diegocheca/goPosts.git/handlers"
	"github.com/gofiber/fiber/v2"
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

	app.Get("/sendingemail", handlers.SendEmail)
	app.Get("/notification/seeder", handlers.NotificationSeeder)
	app.Get("/notification/:id", handlers.ShowNotification)
	app.Get("/notificationshowall", handlers.ShowAllNotification)
	app.Post("/notification", handlers.CreateNotification)

	app.Post("/log", handlers.CreateLog)
	app.Get("/log/show-all", handlers.ShowAllLogs)
	app.Get("/log/show-all-ndjson", handlers.ShowAllLogsNDJson)
	app.Get("/log/seeder", handlers.LogSeeder)
	app.Get("/log/:id", handlers.ShowLog)
	app.Get("/log-telegram", handlers.SendToTelegram)

	app.Get("/emails", handlers.ShowAllEmails)
	app.Get("/emails/:id", handlers.ShowEmail)
	app.Get("/email-seeder", handlers.EmailSeeder)
	app.Post("/emails", handlers.SaveAndSendEmail)
	app.Post("/elk", handlers.CreateELKPost)

}
