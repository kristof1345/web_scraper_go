package scraper

import (
	"testing"
)

func TestScrape(t *testing.T) {
	_, err := ScrapeWeb()
	if err != nil {
		t.Fatal("Test Failed")
	}
}
