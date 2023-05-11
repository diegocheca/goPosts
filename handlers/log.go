package handlers
import (

    "github.com/gofiber/fiber/v2"
    "github.com/diegocheca/goPosts.git/models"
    "github.com/diegocheca/goPosts.git/database"
    "fmt"
    "math/rand"

)

func CreateLog(c *fiber.Ctx) error {
    log := new(models.Log)
    if err := c.BodyParser(log); err != nil {
        return c.Status(500).JSON("error en el body del log que se desea escribir")
    }
    result := database.DB.Db.Create(&log)
    if result.Error != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message":  "Could not create the log", "data": result.Error}) 
    }
    return c.Status(200).JSON(log)
}




func ShowAllLogs(c *fiber.Ctx) error {
	fmt.Println("entre al show all logs")
	
	WriteNotification(c,"entre al showallnotification")
	logs := []models.Log{}
	database.DB.Db.Find(&logs)
    return c.Status(200).JSON(logs)
}







func ShowLog(c *fiber.Ctx) error {
    noty := models.Notification{}
    id := c.Params("id")


    result := database.DB.Db.Where("id = ?", id).First(&noty)
    if result.Error != nil {
        return NotificationNotFound(c)
    }

    return c.Status(200).JSON(noty)
} 


func LogNotFound(c *fiber.Ctx) error {
    //return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
    return c.Status(fiber.StatusNotFound).JSON("notification not found")
}





func LogSeeder(c *fiber.Ctx) error{
	NDATA := 10
	for i := 0; i < NDATA; i++ {

		log := new(models.Log)
        log.UserID = rand.Intn(300)
		log.Controller = "controller"
        log.Function = "function"
		log.Result = "success"
        log.Time = "2022-05-09 12:32:21"
        result := database.DB.Db.Create(&log)
        if result.Error != nil {
            return  c.Status(500).JSON("error")
        }
	}
    return c.Status(200).JSON("log seeder run successfully")
}
