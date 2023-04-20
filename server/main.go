package main

import (
	"encoding/json"
	"fmt"
	"log"

	"example/scraper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

type Rhinos struct {
	// ID          string
	Description string
}

func main() {
	app := fiber.New()

	chanRes := make(chan []byte)

	var rhinos []Rhinos

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/get", func(c *fiber.Ctx) error {
		go scraper.ScrapeWeb(chanRes)

		scrapedData := <-chanRes

		err := json.Unmarshal(scrapedData, &rhinos)
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println(rhinos)

		return c.JSON(rhinos)
	})

	log.Fatal(app.Listen(":4000"))
}
