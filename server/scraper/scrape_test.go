package scraper

import (
	"encoding/json"
	"testing"
)

// func TestScrape(t *testing.T) {
// 	chanRes := make(chan []byte)
// 	go ScrapeWeb(chanRes)
// 	scrapedData := <-chanRes
// 	fmt.Println(scrapedData)
// }

func TestScrapeWeb(t *testing.T) {
	// Create a channel for the output
	out := make(chan []byte)

	// Start the scraping in a separate goroutine
	go ScrapeWeb(out)

	// Receive the scraped data from the output channel
	result := <-out

	// Unmarshal the JSON data into a slice of Fact structs
	var facts []Fact
	err := json.Unmarshal(result, &facts)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	// Check if the length of the facts slice is greater than 0
	if len(facts) == 0 {
		t.Errorf("Expected non-empty facts slice, got empty slice")
	}

	// Check if each fact has a non-empty ID and Description
	for _, fact := range facts {
		if fact.ID == "" || fact.Description == "" {
			t.Errorf("Expected non-empty ID and Description, got empty values")
		}
	}
}

func TestConvertToJson(t *testing.T) {
	// Create some sample data
	facts := []Fact{
		{ID: "1", Description: "Fact 1"},
		{ID: "2", Description: "Fact 2"},
	}

	// Convert the data to JSON
	jsonData := convertToJson(facts)

	// Check if the JSON data is not nil
	if jsonData == nil {
		t.Errorf("Expected non-nil JSON data, got nil")
	}

	// Check if the JSON data can be successfully unmarshalled into a slice of Fact structs
	var unmarshalledFacts []Fact
	err := json.Unmarshal(jsonData, &unmarshalledFacts)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	// Check if the length of the unmarshalled facts slice matches the original slice
	if len(unmarshalledFacts) != len(facts) {
		t.Errorf("Expected %d facts, got %d", len(facts), len(unmarshalledFacts))
	}

	// Check if the unmarshalled facts slice matches the original slice
	for i, fact := range unmarshalledFacts {
		if fact.ID != facts[i].ID || fact.Description != facts[i].Description {
			t.Errorf("Expected Fact{ID: %s, Description: %s}, got Fact{ID: %s, Description: %s}", facts[i].ID, facts[i].Description, fact.ID, fact.Description)
		}
	}
}
