
package main

import (

    "github.com/gofiber/fiber/v2"
	"github.com/diegocheca/goPosts.git/handlers"

	"github.com/diegocheca/goPosts.git/database"
	"github.com/gofiber/template/html"
)

func main() {

	database.ConnectDb()

	engine := html.New("./views", ".html")


    app := fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")

	app.Use(handlers.NotFound)
	
    app.Listen(":3000")
}