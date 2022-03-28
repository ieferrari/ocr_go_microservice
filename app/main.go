package main

import (
	//_ "app/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
	// "github.com/swaggo/fiber-swagger"
	"log"
	// "net/http"
	// "github.com/otiai10/gosseract/v2"
	// "io/ioutil"
	// "strings"
	"runtime"
)

// @Title Ping
// @Summary Will return OK if server is running
// @Accept json
// @Success 200 
// @Router / [get]
func ping(c *fiber.Ctx) error {
	return c.SendString("OK")
} // curl http://localhost:5005/

type statusStruct struct {
	LastLog string `json:"last_log"`
	IsBusy  bool   `json:"is_busy"`
}

var LastLog = ""
var IsBusy = false

// @Title Check Status
// @Summary Willreturn LastLog and IsBusy (true when working, false when is Ready to Work!)
// @Tags Twitter SEARCH
// @Accept json
// @Produce json
// @Success 200 {object} statusStruct{}
// @Failure 503 {object} statusStruct{}
// @Router /check_status [get]
func checkStatus(c *fiber.Ctx) error {
	status := statusStruct{
		LastLog: LastLog,
		IsBusy:  IsBusy,
	}
	return c.JSON(status)
} // curl http://localhost:5005/check_status


// Params is a simple struct for json communication
type Params struct {
	Msg    string `json:"msg"`
}

// @Title OCR for URL
// @Summary Busca una lista de keywords, con uno o varios index, en un periodo  de tiempo
// @Description  keywords es una lista de palabras clave a buscar (puede ser una sola)
// @Tags Twitter SEARCH
// @Accept json
// @Produce json
// @Param   search  body      main.Params  false  "search"
// @Success 200 {object} main.Params{}
// @Failure 503 {object} main.Params{}
// @Router /search [post]
func OcrFromURL(c *fiber.Ctx) error {
	p := Params{}
	if err := c.BodyParser(&p); err != nil {
		return err
	}
	p.Msg = TextFromURL(p.Msg, "spa")
	return c.JSON(p)
} // curl -X POST -d '{"msg": "https://pbs.twimg.com/media/EH-Pvo9WwAEKFwc?format=jpg&name=small"}' -H "Content-Type: application/json" http://127.0.0.1:5005/ocr_from_url

func loadConf() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// if os.Getenv("INGEST_URL") != "" {
	// 	twitter_rd.IngestURL = os.Getenv("INGEST_URL")
	// }
}

// TextFromUrl recibes an url to download an image, converts it to bytes[] and
// send it to gosseract.SetImageFromBytes to extraxT the text.
// You can select the language for the extract: "spa":español, "en":english
// Then, some minor corrections are made before returning the ocr text string
// func TextFromURL(url string, language string)string{
// 	response, e := http.Get(url)
// 	if e != nil {
// 		log.Fatal(e)
// 	}
// 	defer response.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(response.Body)

// 	client := gosseract.NewClient()
// 	defer client.Close()
// 	client.SetImageFromBytes(bodyBytes)
// 	client.SetLanguage(language)
// 	text, _ := client.Text()
// 	text = strings.Replace(text,"º", "o",  -1)
// 	text = strings.Replace(text,"©", "@",  -1)
// 	text = strings.Replace(text,"\n", " ",  -1)
// 	return text	
// }

// @title RD  Twitter-go
// @version 0.1
// @description Twitter client to scrape data, check geoAPI location, and ingest to Logstash
// @termsOfService
// @contact.name support:ieferrari
// @contact.email ivanferrarigalizia@gmail.com
// @host 66.228.41.143:5005
// @BasePath /
func main() {
//	loadConf()
	runtime.GOMAXPROCS(4) // 4 threads/childs max

	app := fiber.New()

	// log.Print(TextFromURL("https://pbs.twimg.com/media/EH-Pvo9WwAEKFwc?format=jpg&name=small", "spa"))

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Get("/", ping)
	app.Get("/ping", ping)
	app.Get("/check_status", checkStatus)
	app.Post("/ocr_from_url", OcrFromURL)
	app.Get("/dashboard", monitor.New())

	log.Println("Ready to work!")
	app.Listen(":5005")
}