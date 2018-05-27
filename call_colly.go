package main

import (
	"github.com/gocolly/colly"
	"fmt"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("a[href]",func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visting",r.URL)
	})
	c.Visit("http://go-colly.org/")
	
}
