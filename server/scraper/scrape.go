package scraper

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

type Fact struct {
	// ID          string `json:"id"`
	Description string `json:"description"`
}

func ScrapeWeb(out chan<- []byte) {
	allFacts := make([]Fact, 0)
	urls := []string{"https://www.zdnet.com/topic/developer/", "https://www.wired.com/tag/developers/"}

	collector := colly.NewCollector(
		colly.AllowedDomains("www.zdnet.com", "zdnet.com", "www.wired.com", "wired.com"),
		// colly.AllowedDomains("www.wired.com", "wired.com"),
	)

	collector.OnHTML("li.item > a", func(element *colly.HTMLElement) {
		factDesc := element.Text

		fact := Fact{
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnHTML(".SummaryItemHedBase-eaxFWE", func(element *colly.HTMLElement) {
		factDesc := element.Text

		fact := Fact{
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	// collector.Visit("https://www.zdnet.com/topic/developer/")
	// collector.Visit("https://www.wired.com/tag/developers/")

	for _, url := range urls {
		collector.Visit(url)
	}

	out <- convertToJson(allFacts)
}

func convertToJson(data []Fact) []byte {
	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return jsonData
}
