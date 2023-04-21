package handlers

import (

    "github.com/gofiber/fiber/v2"
    "github.com/diegocheca/goPosts.git/models"
    "github.com/diegocheca/goPosts.git/database"
    "math/rand"
)



func CreateComment(c *fiber.Ctx) error {
    comment := new(models.Comment)
    if err := c.BodyParser(comment); err != nil {
        return c.Status(500).JSON("error en el body del comment")
    }

    result := database.DB.Db.Create(&comment)
    if result.Error != nil {
        return c.Status(500).JSON("error al escribir el comment en la bd") // NewPostView(c)
    }

    return c.Status(200).JSON(comment)
}

func ShowComment(c *fiber.Ctx) error {
    comment := models.Comment{}
    id := c.Params("id")

    result := database.DB.Db.Where("id = ?", id).First(&comment)
    if result.Error != nil {
        return CommentNotFound(c)
    }

    return c.Status(200).JSON(result)
} 


func CommentNotFound(c *fiber.Ctx) error {
    //return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
    return c.Status(fiber.StatusNotFound).JSON("me fui al 404 de comment")
}


/*
func EditPost(c *fiber.Ctx) error {
    fact := models.Fact{}
    id := c.Params("id")

    result := database.DB.Db.Where("id = ?", id).First(&fact)
    if result.Error != nil {
        return CommentNotFound(c)
    }

    return c.Render("edit", fiber.Map{
        "Title":    "Edit Post",
        "Subtitle": "Edit your interesting Post",
        "Fact":     fact,
    })
}

*/

func UpdateComment(c *fiber.Ctx) error {
    comment := new(models.Comment)
    id := c.Params("id")

    // Parsing the request body
    if err := c.BodyParser(comment); err != nil {
        return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
    }

    // Write updated values to the database
    result := database.DB.Db.Model(&comment).Where("id = ?", id).Updates(comment)
    if result.Error != nil {
        return EditFact(c)
    }

    return c.Status(200).JSON(comment)
}


func DeleteComment(c *fiber.Ctx) error {
    comment := models.Comment{}
    id := c.Params("id")
    result := database.DB.Db.Where("id = ?", id).Delete(&comment)
    if result.Error != nil {
        return CommentNotFound(c)
    }
    return c.Status(200).JSON("Comment eliminado correctamente")
}






func CommentSeeder(c *fiber.Ctx) error{
	NDATA := 800
	for i := 0; i < NDATA; i++ {
		comment := new(models.Comment)
        comment.PostID = rand.Intn(800)
        comment.Content = randomContent()
        comment.Image = "https://source.unsplash.com/random/?food"
        comment.Author = rand.Intn(100)
        comment.Likes = rand.Intn(10000)
        result := database.DB.Db.Create(&comment)
        if result.Error != nil {
            return  c.Status(500).JSON("error")
        }
	}
    return c.Status(200).JSON("comment seeder run successfully")
}