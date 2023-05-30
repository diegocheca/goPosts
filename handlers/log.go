package handlers

import (
	"github.com/diegocheca/goPosts.git/database"
	"github.com/diegocheca/goPosts.git/models"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	//"github.com/slack-go/slack"

	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var listOfControllers = []string{"Blog", "Route", "Place", "State", "Comment", "Like", "User", "Path", "TravelGuisde", "Notifications", "Logs"}
var listOfFunctions = []string{"Index", "Edit", "Update", "Show", "ShowAll", "Destroy"}
var listOfResults = []string{"success", "error", "500", "404", "501", "402"}
var listOfBrowser = []string{"Google Chrome", "Mozilla Firefox", "Opera", "Apple Safari", "Microsoft Edge", "Netscape", ""}
var listOfOs = []string{"Windows", "Linux", "Apple", "IOS", "Android", ""}
var listOfApp = []string{"Mobile v1.1", "Mobile v1.2", "Mobile v1.3", "Mobile v2.0", "Mobile v2.1", "Mobile v3.0", "Web"}
var listOfTable = []string{"User", "Profile", "Blogs", "Comments", "Likes", "Images", "Activities", "Routes", "States", "Places", "TravelGuides", "Paths", "Logs", "Emails", "Notifications", "Monitoring"}
var listOfMicros = []string{"Blogs", "Gateway", "Authorization", "Routes", "GO", "ELK", "Jenkins", "GitLab", "MailHog", "Redis", "Postgres", "MySql"}

const charset = "abcdefghijklmnopqrstuvwxyz" + " " + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

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
	req.SetRequestURI("https://api.telegram.org/bot" + os.Getenv("TELEGRAM_API") + "/sendMessage?chat_id=1485456302&text=Hello%20Bro%20Im%20gofer22")
	return client.Do(req, res)

}
