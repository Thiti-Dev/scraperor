package core

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type ScraperorEngine struct {
	collector *colly.Collector
	target_url string
	pointer string
}

func (engine *ScraperorEngine) AddTask(url string,pointer string){
	engine.target_url = url
	engine.pointer = pointer
}

// Required AddTask to be excecuted before this
func (engine *ScraperorEngine) Scrape() []string{
	produced_elements := make([]string,0)
	engine.collector.OnHTML(engine.pointer, func(e *colly.HTMLElement) {
		//fmt.Println(e.Text)
		produced_elements = append(produced_elements, e.Text)
	})
	engine.collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	engine.collector.Visit(engine.target_url)
	return produced_elements
}

func CreateEngineEntity() *ScraperorEngine {
	return &ScraperorEngine{
		collector: colly.NewCollector(),
	}
}