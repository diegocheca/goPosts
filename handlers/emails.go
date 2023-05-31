package handlers
import (

    "github.com/gofiber/fiber/v2"
    "github.com/diegocheca/goPosts.git/models"
    "github.com/diegocheca/goPosts.git/database"
    "math/rand"

    "fmt"



)

var listOfEnviorment = []string{"local", "prod", "preprod", "staging"}

func CreateEmail(c *fiber.Ctx) error {
    email := new(models.Emails)
    if err := c.BodyParser(email); err != nil {
        return c.Status(500).JSON("error en el body del email")
    }
    result := database.DB.Db.Create(&email)
    if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message":  "Could not create email", "data": result.Error}) 
    }

    return c.Status(200).JSON(email)
}




func ShowAllEmails(c *fiber.Ctx) error {
	fmt.Println("entre al show all emails")
	
	WriteNotification(c,"entre al ShowAllEmails")
	emails := []models.Emails{}
	database.DB.Db.Find(&emails)
    return c.Status(200).JSON(emails)
}




func ShowEmail(c *fiber.Ctx) error {
    email := models.Emails{}
    id := c.Params("id")


    result := database.DB.Db.Where("id = ?", id).First(&email)
    if result.Error != nil {
        return NotificationNotFound(c)
    }

    return c.Status(200).JSON(email)
} 


func EmailNotFound(c *fiber.Ctx) error {
    //return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
    return c.Status(fiber.StatusNotFound).JSON("email not found")
}



func EmailSeeder(c *fiber.Ctx) error{
	NDATA := 800
	for i := 0; i < NDATA; i++ {
		randomDate := randate()
		newEmail := ""
		newEmail = StringRan(10)+"@"+StringRan(5)+".com"
		email := new(models.Emails)
        email.UserID = rand.Intn(300)
		email.EmailTo = newEmail
		newEmail = StringRan(10)+"@"+StringRan(5)+".com"
        email.EmailFrom = newEmail
		email.Enviornment = listOfEnviorment[rand.Intn(len(listOfEnviorment))]
        email.Subject = randomContent()
        email.Body = randomContent()
        email.Result = listOfResults[rand.Intn(len(listOfResults))]
		email.Time = randomDate.Format("2006-01-02 15:04:05.000000")
        result := database.DB.Db.Create(&email)
        if result.Error != nil {
            return  c.Status(500).JSON("error")
        }
	}
    return c.Status(200).JSON("notification seeder run successfully")
}

