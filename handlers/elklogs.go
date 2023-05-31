package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"

	log "github.com/sirupsen/logrus"
)

func CreateELKPost(c *fiber.Ctx) error {

	//return c.Status(200).JSON("termine de correr la funcion de elk")

	f, err := os.OpenFile("example.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON("error al escribir en el log")
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
	log.WithFields(log.Fields{
		"event":   "create_profile",
		"user_id": 10,
	}).Info("This is an info message.")

	log.WithFields(log.Fields{
		"event":   "delete_profile",
		"user_id": 11,
	}).Warn("This is a warning message.")

	log.WithFields(log.Fields{
		"event":   "edit_profile",
		"user_id": 13,
		"package": "main",
	}).Fatal("This is a critical message.")

	return c.Status(200).JSON("termine de correr la funcion de elk")
}
