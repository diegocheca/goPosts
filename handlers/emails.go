package handlers

import (
	"math/rand"

	"github.com/diegocheca/goPosts.git/database"
	"github.com/diegocheca/goPosts.git/models"
	"github.com/gofiber/fiber/v2"

	"fmt"
	"net/smtp"
	"os"
	"strconv"

	"github.com/jordan-wright/email"
)

/*
Sure, here is a description for a Golang handler called EmailsHandler that sends emails with some notifications to clients in a social media system:

The EmailsHandler handler is responsible for sending emails with notifications to clients in a social media system. It provides methods for sending emails with different types of notifications, such as new friend requests, new messages, and new posts. The handler also provides methods for configuring the email settings, such as the email server, the email address, and the email password.

The following are the methods that are provided by the EmailsHandler handler:

sendEmail(): This method sends an email with a notification to a client.
sendNewFriendRequestEmail(): This method sends an email with a notification to a client about a new friend request.
sendNewMessageEmail(): This method sends an email with a notification to a client about a new message.
sendNewPostEmail(): This method sends an email with a notification to a client about a new post.
configureEmailSettings(): This method configures the email settings, such as the email server, the email address, and the email password.
The EmailsHandler handler is used by the NotificationService to send emails with notifications to clients. The NotificationService is a service that provides access to notifications in a social media system.

Here are some additional details about the EmailsHandler handler:
*/
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

func SaveAndSendEmail(c *fiber.Ctx) error {

	emailToDb := new(models.Emails)
	if err := c.BodyParser(emailToDb); err != nil {
		return c.Status(500).JSON("error en los parametros del email")
	}

	result := database.DB.Db.Create(&emailToDb)
	if result.Error != nil {
		return c.Status(500).JSON("error al guardar email en la bd")
	}

	port, err := strconv.Atoi(os.Getenv("SERVER_EMAIL_PORT"))
	if err != nil {
		return c.Status(500).JSON("error al acceder al puerto del server de emails")
	}

	m := struct {
		Host string
		Port int
		User string
		Pass string
	}{

		Host: os.Getenv("SERVER_EMAIL_IP"), // resolved to 127.0.0.1 in /etc/hosts
		Port: port,
		User: ``,
		Pass: ``,
	}
	e := email.NewEmail()
	e.From = emailToDb.EmailFrom
	e.To = []string{emailToDb.EmailTo}
	e.Subject = emailToDb.Subject
	e.Text = []byte(emailToDb.Body)
	errToSend := e.Send(fmt.Sprintf("%s:%d", m.Host, m.Port), smtp.CRAMMD5Auth(m.User, m.Pass))
	if errToSend != nil {
		fmt.Println(errToSend)
		return c.Status(500).JSON(errToSend)
	}
	fmt.Println("Email Sent Successfully!")

	return c.Status(200).JSON("listo")

}

func CreateEmail(c *fiber.Ctx) error {
	email := new(models.Emails)
	if err := c.BodyParser(email); err != nil {
		return c.Status(500).JSON("error en el body del email")
	}
	result := database.DB.Db.Create(&email)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create email", "data": result.Error})
	}
	return c.Status(200).JSON(email)
}

func ShowAllEmails(c *fiber.Ctx) error {
	fmt.Println("entre al show all emails")
	WriteNotification(c, "entre al ShowAllEmails")
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

func EmailSeeder(c *fiber.Ctx) error {
	NDATA := 100
	for i := 0; i < NDATA; i++ {
		randomDate := randate()
		email := new(models.Emails)
		email.UserID = rand.Intn(300)
		email.EmailTo = fakeEmails[rand.Intn(len(fakeEmails))]
		email.EmailFrom = fakeEmails[rand.Intn(len(fakeEmails))]
		email.Enviornment = listOfEnviorment[rand.Intn(len(listOfEnviorment))]
		email.Subject = fakeSubject[rand.Intn(len(fakeSubject))]
		email.Body = fakeEmailBodies[rand.Intn(len(fakeEmailBodies))]
		email.Result = listOfResults[rand.Intn(len(listOfResults))]
		email.Time = randomDate.Format("2006-01-02 15:04:05.000000")
		fmt.Println("Antes de guardar en la bd")
		result := database.DB.Db.Create(&email)
		fmt.Println("Despues de la bd")
		if result.Error != nil {
			return c.Status(500).JSON("error")
		}
	}
	return c.Status(200).JSON("Emails seeder run successfully")
}

/*
welcome email
Subject: Welcome to DeguAPP!

Hi [client name],

Welcome to DeguAPP! We're excited to have you join our community.

DeguAPP is a social media platform that allows you to connect with friends, family, and like-minded people from all over the world. You can share photos, videos, and stories, and connect with people who share your interests.

To get started, simply create a profile and start exploring. You can find friends by searching for them by name, email address, or phone number. You can also join groups and communities that interest you.

We hope you enjoy using DeguAPP! If you have any questions, please don't hesitate to contact us.

Thanks,
The DeguAPP Team




Asunto: ¡Bienvenido a DeguAPP!

Hola [nombre del cliente],

¡Bienvenido a DeguAPP! Estamos emocionados de que te unas a nuestra comunidad.

DeguAPP es una plataforma de redes sociales que le permite conectarse con amigos, familiares y personas de ideas afines de todo el mundo. Puede compartir fotos, videos e historias y conectarse con personas que comparten sus intereses.

Para comenzar, simplemente cree un perfil y comience a explorar. Puede encontrar amigos buscándolos por nombre, dirección de correo electrónico o número de teléfono. También puedes unirte a grupos y comunidades que te interesen.

¡Esperamos que disfrute usando DeguAPP! Si tiene alguna pregunta, no dude en ponerse en contacto con nosotros.

Gracias,
El equipo de DeguAPP



*/
