package myService

import (
	"fmt"

	"github.com/gocolly/colly"
)

//爬虫。
func CrawlerStar(url string) {
	c := colly.NewCollector()
	print("-----------------d---", url)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Println(link)
		// Visit link found on page
		e.Request.Visit(link)
	})

	// Start scraping on https://en.wikipedia.org
	c.Visit(url)
}
