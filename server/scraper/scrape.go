package scraper

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
)

type Fact struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

func ScrapeWeb(out chan<- []byte, s []string) {
	allFacts := make([]Fact, 0)
	currURL := ""
	// urls := []string{"https://www.zdnet.com/topic/developer/"}

	collector := colly.NewCollector(
	// colly.Async(true),
	// colly.AllowedDomains("www.zdnet.com", "zdnet.com", "www.wired.com", "wired.com"),
	// colly.AllowedDomains("www.wired.com", "wired.com"),
	)

	collector.OnHTML("li.item > a", func(element *colly.HTMLElement) {
		u, err := url.Parse(currURL)
		if err != nil {
			panic(err)
		}
		u.Path = ""
		u.RawQuery = ""
		u.Fragment = ""

		factDesc := element.Text
		urlPart := element.Attr("href")
		urlFull := u.String() + urlPart

		fact := Fact{
			URL:         urlFull,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnHTML(".SummaryItemHedLink-ciaMYZ", func(element *colly.HTMLElement) {
		u, err := url.Parse(currURL)
		if err != nil {
			panic(err)
		}
		u.Path = ""
		u.RawQuery = ""
		u.Fragment = ""

		factDesc := element.ChildText("h3")
		urlPart := element.Attr("href")
		urlFull := u.String() + urlPart

		fact := Fact{
			URL:         urlFull,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	// collector.Visit("https://www.zdnet.com/topic/developer/")
	// collector.Visit("https://www.wired.com/tag/developers/")

	for _, url := range s {
		currURL = url
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
