package handlers

import (
	"fmt"
	"os"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/diegocheca/goPosts.git/database"
	"github.com/diegocheca/goPosts.git/models"
	"github.com/gofiber/fiber/v2"
	logFile "github.com/sirupsen/logrus"
)

type LogDos struct {
	UserID      int
	Controller  string
	Function    string
	Result      string
	Browser     string
	Os          string
	App         string
	Ip          string
	Table       string
	TableId     int
	Action      string
	Micro       string
	Description string
	Time        string
}

func (logdoslocal LogDos) Validate() error {
	return validation.ValidateStruct(
		&logdoslocal,
		validation.Field(&logdoslocal.UserID, validation.Required),
		validation.Field(&logdoslocal.Controller, validation.Required),
		validation.Field(&logdoslocal.Function, validation.Required),
		validation.Field(&logdoslocal.Result, validation.Required),
		validation.Field(&logdoslocal.Browser, validation.Required),
		validation.Field(&logdoslocal.Os, validation.Required),
		validation.Field(&logdoslocal.App, validation.Required),

		validation.Field(&logdoslocal.Ip, validation.Required),
		validation.Field(&logdoslocal.Table, validation.Required),
		validation.Field(&logdoslocal.TableId, validation.Required),
		validation.Field(&logdoslocal.Action, validation.Required),
		validation.Field(&logdoslocal.Micro, validation.Required),
		validation.Field(&logdoslocal.Description, validation.Required),
		validation.Field(&logdoslocal.Time, validation.Required),
	)

}
func CreateELKPost(c *fiber.Ctx) error {

	my_log := new(models.Log)
	if err := c.BodyParser(my_log); err != nil {
		return c.Status(500).JSON("error en el body del log que se desea escribir")
	}

	//luego de crear el ob lo valido
	/*data := "example"
	err := validation.Validate(data,
		validation.Required,       // not empty
		validation.Length(5, 100), // length between 5 and 100
		//is.URL,                    // is a valid URL
	)
	fmt.Println(err)*/
	logdoslocal := LogDos{
		UserID:      4,
		Controller:  "Controller",
		Function:    "Function",
		Result:      "Result",
		Browser:     "Browser",
		Os:          "Os",
		App:         "App",
		Ip:          "Ip",
		Table:       "Table",
		TableId:     5,
		Action:      "Action",
		Micro:       "Micro",
		Description: "Description",
		Time:        "Time",
	}
	err := logdoslocal.Validate()
	fmt.Println(err)
	//una vez valido lo escribo en la bd

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
		"user_id":     my_log.UserID,
		"controller":  my_log.Controller,
		"function":    my_log.Function,
		"result":      my_log.Result,
		"browser":     my_log.Browser,
		"os":          my_log.Os,
		"app":         my_log.App,
		"ip":          my_log.Ip,
		"table":       my_log.Table,
		"table_id":    my_log.TableId,
		"action":      my_log.Action,
		"micro":       my_log.Micro,
		"description": my_log.Description,
		"time":        my_log.Time,
	}).Info("This is an info message.")

	/*log.WithFields(log.Fields{
		"event":   "delete_profile",
		"user_id": 11,
	}).Warn("This is a warning message.")*/

	defer f.Close()

	return c.Status(200).JSON("\"status\":\"success\",\"message\":\"Log written successfully\"")

	return c.Status(200).JSON("log escrito correctamente")

	return c.Status(200).JSON("log escrito correctamente")
	defer f.Close()
	return c.Status(200).JSON("termine de correr la funcion de elk")
}
