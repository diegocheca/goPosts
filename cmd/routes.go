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

	//POSTS
	app.Get("/posts", handlers.ListPosts)
	//app.Get("/post/new", handlers.NewPostView)
	app.Post("/post", handlers.CreatePost)
	app.Get("/post/:id", handlers.ShowPost)
	//app.Get("/post/:id/edit", handlers.EditPost)
	app.Patch("/post/:id", handlers.UpdatePost)
	app.Delete("/post/:id", handlers.DeletePost)
	app.Post("/post/seeder", handlers.PostSeeder)

	//Comments
	app.Get("/comments", handlers.ListComments)
	//app.Get("/post/new", handlers.NewPostView)
	app.Post("/comment", handlers.CreateComment)
	app.Get("/comment/:id", handlers.ShowComment)
	//app.Get("/comment/:id/edit", handlers.EditComment)
	app.Patch("/comment/:id", handlers.UpdateComment)
	app.Delete("/comment/:id", handlers.DeleteComment)
	app.Post("/comment/seeder", handlers.CommentSeeder)

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
	app.Post("/emails", handlers.CreateEmail)
    app.Get("/emails/:id", handlers.ShowEmail)
	app.Get("/emails/seeder", handlers.EmailSeeder)
	
}
