package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "artists.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) { fmt.Println("Scraping:", r.URL) })
	c.OnResponse(func(r *colly.Response) { fmt.Println("Status:", r.StatusCode) })
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "nError:", err)
	})

	// c.OnHTML("*", func(e *colly.HTMLElement) {
	// 	fmt.Println(e)
	// })

	// c.OnHTML("a.sc-dIvrsQ", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	c.OnHTML(("div.sc-jYKCQm.fidnen"), func(e *colly.HTMLElement) {
		fmt.Println("here")
	})
	c.Visit("https://www.stubhub.com/taylor-swift-philadelphia-tickets-5-12-2023/event/150593637/?quantity=2")

}
