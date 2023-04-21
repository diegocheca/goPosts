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
    //app.Get("/comments", handlers.ListPosts)
    //app.Get("/post/new", handlers.NewPostView)
    app.Post("/comment", handlers.CreateComment)
    app.Get("/comment/:id", handlers.ShowComment)
    //app.Get("/comment/:id/edit", handlers.EditComment)
    app.Patch("/comment/:id", handlers.UpdateComment)
    app.Delete("/comment/:id", handlers.DeleteComment)
    app.Post("/comment/seeder", handlers.CommentSeeder)

}