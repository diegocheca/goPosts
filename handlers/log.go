package handlers

import (
	"github.com/diegocheca/goPosts.git/database"
	"github.com/diegocheca/goPosts.git/models"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	//"github.com/slack-go/slack"
	"net/url"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + " " + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

/*
The logHandler handler is responsible for performing CRUD operations on logs in a social media system.
It provides methods for creating, reading, updating, and deleting logs. The handler also provides methods for filtering
logs by user, type, and status.
The logHandler handler uses the log package to write logs to a file.
The logHandler handler uses the context package to pass information about the current request to the LogService.
The logHandler handler uses the errors package to handle errors that occur when performing CRUD operations on logs.

createLog(): This method creates a new log.
readLogs(): This method returns all of the logs for a given user.
updateLog(): This method updates an existing log.
deleteLog(): This method deletes a log.
filterLogs(): This method returns a filtered list of logs.
*/
func CreateLog(c *fiber.Ctx) error {
	log := new(models.Log)
	if err := c.BodyParser(log); err != nil {
		return c.Status(500).JSON("error en el body del log que se desea escribir")
	}
	result := database.DB.Db.Create(&log)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create the log", "data": result.Error})
	}
	return c.Status(200).JSON(log)
}

func ShowAllLogs(c *fiber.Ctx) error {
	fmt.Println("entre al show all logs")

	WriteNotification(c, "entre al showallnotification")
	logs := []models.Log{}
	database.DB.Db.Find(&logs)
	return c.Status(200).JSON(logs)
}

func ShowAllLogsNDJson(c *fiber.Ctx) error {

	fmt.Println("entre al show all logs con NDJson")

	WriteNotification(c, "entre al showallnotification con NDJson")
	logs := []models.Log{}
	database.DB.Db.Find(&logs)

	jsonData, err := json.Marshal(logs)

	if err != nil {
		//fmt.Printf("could not marshal json: %s\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create the log", "data": err})
	}

	dst := &bytes.Buffer{}
	if err := json.Compact(dst, []byte(jsonData)); err != nil {
		panic(err)
	}
	//fmt.Println(dst.String())

	return c.Status(200).JSON(dst.String())

}

func ShowLog(c *fiber.Ctx) error {
	log := models.Log{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&log)
	if result.Error != nil {
		return NotificationNotFound(c)
	}

	return c.Status(200).JSON(log)
}

func LogNotFound(c *fiber.Ctx) error {
	//return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
	return c.Status(fiber.StatusNotFound).JSON("notification not found")
}

func RandomIp() string {
	return strconv.Itoa(rand.Intn(255)) + "." + strconv.Itoa(rand.Intn(255)) + "." + strconv.Itoa(rand.Intn(255)) + "." + strconv.Itoa(rand.Intn(255))
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func StringRan(length int) string {
	return StringWithCharset(length, charset)
}

func randate() time.Time {
	min := time.Date(2018, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

// install
//
//	go get -u github.com/valyala/fasthttp
func LogSeeder(c *fiber.Ctx) error {
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
	NDATA := 1000
	current := time.Now()
	message := ""
	for i := 0; i < NDATA; i++ {

		randomDate := randate()
		//fechaComoTime, err := time.Parse("2006-01-02 15:04:05.000000", randomDate)

		mylog := new(models.Log)
		mylog.UserID = rand.Intn(300)
		mylog.Controller = listOfControllers[rand.Intn(len(listOfControllers))]
		mylog.Function = listOfFunctions[rand.Intn(len(listOfFunctions))]
		mylog.Result = listOfResults[rand.Intn(len(listOfResults))]
		mylog.Time = randomDate.Format("2006-01-02 15:04:05.000000")

		mylog.Browser = listOfBrowser[rand.Intn(len(listOfBrowser))]
		mylog.Os = listOfOs[rand.Intn(len(listOfOs))]
		mylog.App = listOfApp[rand.Intn(len(listOfApp))]

		mylog.Ip = RandomIp()
		mylog.Table = listOfTable[rand.Intn(len(listOfTable))]

		mylog.TableId = rand.Intn(999999)
		mylog.Action = listOfFunctions[rand.Intn(len(listOfFunctions))]

		mylog.Micro = listOfMicros[rand.Intn(len(listOfMicros))]
		mylog.Description = StringRan(100)

		result := database.DB.Db.Create(&mylog)
		message = strconv.Itoa(mylog.UserID) + " - " + mylog.Controller + " - " + mylog.Function + " - " + mylog.Result + " - " + current.Format("2006-01-02 15:04:05.000000")
		log.Printf("Log:  %s", message)
		if result.Error != nil {
			return c.Status(500).JSON("error")
		}
		//time.Sleep(1 * time.Second)

	}
	return c.Status(200).JSON("log seeder run successfully")
}

func SendToTelegram(c *fiber.Ctx) error {
	client := fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		DisablePathNormalizing:   true,
	}
	req := c.Request()
	res := c.Response()
	//os.Getenv("TELEGRAM_API")
	type message struct{
		Mensaje string `json: "mensaje"`
	}
	my_message := new(message)
    if err := c.BodyParser(my_message); err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could get the telegram url for make the request", "data": err.Error})
    }


	baseUrl, err := url.Parse(os.Getenv("TELEGRAM_URL")+ os.Getenv("TELEGRAM_API"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could get the telegram url for make the request", "data": err.Error})
	
	}

	// Add a Path Segment (Path segment is automatically escaped)
	baseUrl.Path += "/sendMessage"

	// Prepare Query Parameters
	params := url.Values{}
	params.Add("chat_id", os.Getenv("TELEGRAM_CHAT_ID"))
	params.Add("text", my_message.Mensaje)

	// Add Query Parameters to the URL
	baseUrl.RawQuery = params.Encode() // Escape Query Parameters

	req.SetRequestURI(baseUrl.String())
	return client.Do(req, res)

}
