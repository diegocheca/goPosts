package handlers
import (

    "github.com/gofiber/fiber/v2"
    "github.com/diegocheca/goPosts.git/models"
    "github.com/diegocheca/goPosts.git/database"
    "github.com/valyala/fasthttp"

    //"github.com/slack-go/slack"

    "fmt"
    "math/rand"
    "time"
    "log"
	"os"
    "strconv"
)
var listOfControllers = []string{"Blog", "Route","Place", "State","Comment", "Like", "User","Path","TravelGuisde","Notifications","Logs"}
var listOfFunctions = []string {"Index", "Edit","Update", "Show","ShowAll", "Destroy"}
var listOfResults = []string {"success", "error","500", "404","501", "402"}

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


//install
//   go get -u github.com/valyala/fasthttp
func LogSeeder(c *fiber.Ctx) error{
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
	NDATA := 10
    current := time.Now()
    message := ""
	for i := 0; i < NDATA; i++ {
		mylog := new(models.Log)
        mylog.UserID = rand.Intn(300)
		mylog.Controller = listOfControllers[rand.Intn(len(listOfControllers))]
        mylog.Function = listOfFunctions[rand.Intn(len(listOfFunctions))]
		mylog.Result = listOfResults[rand.Intn(len(listOfResults))]
        mylog.Time = current.Format("2006-01-02 15:04:05.000000")
        result := database.DB.Db.Create(&mylog)
        message = strconv.Itoa(mylog.UserID)+" - "+mylog.Controller+" - "+mylog.Function+" - "+mylog.Result+" - "+current.Format("2006-01-02 15:04:05.000000")
        log.Printf("Log:  %s", message)
        if result.Error != nil {
            return  c.Status(500).JSON("error")
        }
	}
    return c.Status(200).JSON("log seeder run successfully")
}


func SendToTelegram(c *fiber.Ctx) error{
    client := fasthttp.Client{
        NoDefaultUserAgentHeader: true,
        DisablePathNormalizing:   true,
    }
    req := c.Request()
    res := c.Response()
    //os.Getenv("TELEGRAM_API")
    //req.SetRequestURI("https://api.telegram.org/bot"+os.Getenv("TELEGRAM_API")+"/sendMessage?chat_id=1485456302&text=Hello%20Bro%20Im%20gofer22")
    return client.Do(req, res)

}