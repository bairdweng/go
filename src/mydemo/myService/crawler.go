package myService

import (
	"github.com/gocolly/colly"
)

//爬虫。
func CrawlerStar() {
	c := colly.NewCollector()
	print(c)
}
