package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"

	"github.com/diegocheca/goPosts.git/database"
	"github.com/diegocheca/goPosts.git/models"
	"github.com/gofiber/fiber/v2"

	"github.com/jordan-wright/email"
)

/*
This notificactionController controller is responsible for performing CRUD operations on notifications in a social media system.
It provides methods for creating, reading, updating, and deleting notifications. The controller also provides methods for filtering notifications by user, type, and status.
createNotification(): This method creates a new notification.
readNotifications(): This method returns all of the notifications for a given user.
updateNotification(): This method updates an existing notification.
deleteNotification(): This method deletes a notification.
filterNotifications(): This method returns a filtered list of notifications.
*/
func WriteNotification(c *fiber.Ctx, message string) error {
	// abrir el archivo webserver.log para escritura
	f, err := os.OpenFile("webserver.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON("error al escribir en el log")
	}
	// y cerrar cuando termine la funcion main
	defer f.Close()
	// asociar el manejador del archivo al log
	log.SetOutput(f)
	// agregan 10 lineas al archivo log
	log.Printf("Log:  %s", message)
	return c.Status(200).JSON("log escrito correctamente")
}

func CreateNotification(c *fiber.Ctx) error {
	noty := new(models.Notification)
	if err := c.BodyParser(noty); err != nil {
		return c.Status(500).JSON("error en el body de la notification")
	}
	result := database.DB.Db.Create(&noty)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": result.Error})
	}

	return c.Status(200).JSON(noty)
}

func ShowAllNotification(c *fiber.Ctx) error {
	fmt.Println("entre al showallnotification")

	WriteNotification(c, "entre al showallnotification")
	noty := []models.Notification{}
	database.DB.Db.Find(&noty)
	return c.Status(200).JSON(noty)
}

func SendEmail(c *fiber.Ctx) error {

	m := struct {
		Host string
		Port int
		User string
		Pass string
	}{
		Host: `192.168.1.120`, // resolved to 127.0.0.1 in /etc/hosts
		Port: 1025,
		User: ``,
		Pass: ``,
	}
	e := email.NewEmail()
	e.From = `notifications@degustur.com`
	e.To = []string{`diegochecarelli@hotmail.com`}
	e.Subject = `Go Email`
	e.Text = []byte(`testi2222222ng email from my linux with go`)
	err := e.Send(fmt.Sprintf("%s:%d", m.Host, m.Port), smtp.CRAMMD5Auth(m.User, m.Pass))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(err)
	}
	fmt.Println("Email Sent Successfully!")

	return c.Status(200).JSON("listo")

}

func ShowNotification(c *fiber.Ctx) error {
	noty := models.Notification{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&noty)
	if result.Error != nil {
		return NotificationNotFound(c)
	}

	return c.Status(200).JSON(noty)
}

func NotificationNotFound(c *fiber.Ctx) error {
	//return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
	return c.Status(fiber.StatusNotFound).JSON("notification not found")
}

func NotificationSeeder(c *fiber.Ctx) error {
	NDATA := 800
	for i := 0; i < NDATA; i++ {

		channels := [6]string{"email", "sms", "telegram", "slack", "log", "push"}

		readed := false
		extra_data := ""
		if rand.Intn(10) > 5 {
			readed = true
			extra_data = randomContent()
		}

		noty := new(models.Notification)
		noty.UserID = rand.Intn(300)
		noty.UserIdFrom = rand.Intn(300)
		noty.Readed = readed
		noty.Content = randomContent()
		noty.NotificationType = rand.Intn(5)
		noty.ExtraData = extra_data
		noty.Channel = channels[rand.Intn(len(channels))]
		result := database.DB.Db.Create(&noty)
		if result.Error != nil {
			return c.Status(500).JSON("error")
		}
	}
	return c.Status(200).JSON("notification seeder run successfully")
}
