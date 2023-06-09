package main

import (
	"encoding/json"
	"fmt"
	"log"

	"example/scraper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Data struct {
	URL         string
	Description string
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:8000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setupRoutes(app)

	log.Fatal(app.Listen(":4000"))
}

func setupRoutes(app *fiber.App) {
	// app.Get("/get", returnScrapedData)
	app.Post("/api", returnPostedData)
}

// func returnScrapedData(c *fiber.Ctx) error {
// 	chanRes := make(chan []byte)

// 	s := []string{"https://www.zdnet.com/topic/developer/"}

// 	var rhinos []Rhinos

// 	go scraper.ScrapeWeb(chanRes, s)

// 	scrapedData := <-chanRes

// 	err := json.Unmarshal(scrapedData, &rhinos)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}

// 	return c.JSON(rhinos)
// }

func returnPostedData(c *fiber.Ctx) error {
	URLS := struct {
		Url1  string `json:"url1"`
		ELem1 string `json:"item1"`
		Url2  string `json:"url2"`
		Elem2 string `json:"item2"`
	}{}
	if err := c.BodyParser(&URLS); err != nil {
		return err
	}
	s := []string{
		URLS.Url1,
		URLS.Url2,
	}
	e := []string{
		URLS.ELem1,
		URLS.Elem2,
	}
	chanRes := make(chan []byte)

	var data []Data

	go scraper.ScrapeWeb(chanRes, s, e)

	scrapedData := <-chanRes

	err := json.Unmarshal(scrapedData, &data)
	if err != nil {
		fmt.Println("error:", err)
	}

	return c.JSON(data)
}
