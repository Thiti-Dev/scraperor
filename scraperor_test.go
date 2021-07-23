package main

import (
	"testing"

	"github.com/Thiti-Dev/scraperor/core"
)

// TestScraping creates a colly collector and then scrape the author github bio description
// expects atleast 1 element returned
func TestScraping(t *testing.T) {
	engineEntity := core.CreateEngineEntity()
	engineEntity.AddTask("https://github.com/Thiti-Dev",".user-profile-bio div")
	
	elements := engineEntity.Scrape()
	if len(elements) == 0{
		t.Fatal("[Colly]: having a problem in getting github bio")
	}
}