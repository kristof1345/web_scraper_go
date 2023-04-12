package scraper

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func ScrapeWeb(out chan<- []byte) {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factId := element.Attr("id")

		factDesc := element.Text

		fact := Fact{
			ID:          factId,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")

	out <- convertToJson(allFacts)
}

func convertToJson(data []Fact) []byte {
	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return jsonData
}
