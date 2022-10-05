package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var url string = "https://en.wikipedia.org/wiki/Web_scraping"

func main() {
	scrape(url)
}

func scrape(url string) {
	fmt.Println("Hola!")
	print("hola con print\n")

	c := colly.NewCollector()

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visitando:", request.URL)
	})

	// fmt.Println(c)

	// Find and print all links
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttr("a", "href")
		fmt.Println(links)
	})

	c.Visit(url)

}
