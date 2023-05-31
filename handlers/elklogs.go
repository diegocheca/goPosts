package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/diegocheca/goPosts.git/database"
	"github.com/diegocheca/goPosts.git/models"



	logFile "github.com/sirupsen/logrus"
)

func CreateELKPost(c *fiber.Ctx) error {
	my_log := new(models.Log)
	if err := c.BodyParser(my_log); err != nil {
		return c.Status(500).JSON("error en el body del log que se desea escribir")
	}
	result := database.DB.Db.Create(&my_log)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create the log", "data": result.Error})
	}
	//return c.Status(200).JSON(log)



	f, err := os.OpenFile("example.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		//log.Fatal(err)
		return c.Status(500).JSON("error al escribir en el log")
	}
	// y cerrar cuando termine la funcion main
	defer f.Close()

	logFile.SetFormatter(&logFile.JSONFormatter{})
	logFile.SetOutput(f)
	logFile.WithFields(logFile.Fields{
		"user_id":   my_log.UserID,
		"controller":   my_log.Controller,
		"function":   my_log.Function,
		"result":   my_log.Result,
		"browser":   my_log.Browser,
		"os":   my_log.Os,
		"app":   my_log.App,
		"ip":   my_log.Ip,
		"table":   my_log.Table,
		"table_id":   my_log.TableId,
		"action":   my_log.Action,
		"micro":   my_log.Micro,
		"description":   my_log.Description,
		"time":   my_log.Time,
	}).Info("This is an info message.")


	/*log.WithFields(log.Fields{
		"event":   "delete_profile",
		"user_id": 11,
	}).Warn("This is a warning message.")*/

	
	defer f.Close()
	
	return c.Status(200).JSON("log escrito correctamente")


	
	return c.Status(200).JSON("log escrito correctamente")
	defer f.Close()
	return c.Status(200).JSON("termine de correr la funcion de elk")
}
