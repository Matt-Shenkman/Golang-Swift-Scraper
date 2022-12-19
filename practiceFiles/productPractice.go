package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/gocolly/colly"
// )

// type Product struct {
// 	name            string
// 	rating          string
// 	numberOfRatings int
// 	price           string
// 	retailPrice     string
// }

// func main() {
// 	c := colly.NewCollector(colly.AllowedDomains())
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})

// 	// productList := make([]Product, 0)

// 	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
// 		// var p Product

// 		e.ForEach("div.a-section.a-spacing-small.a-spacing-top-small", func(_ int, e *colly.HTMLElement) {
// 			var productName string
// 			var rating string
// 			var numberOfRatings string
// 			var price string

// 			fName := "amazon.csv"
// 			file, err := os.Create(fName)
// 			if err != nil {
// 				log.Fatalf("Could not create file, err: %q", err)
// 				return
// 			}
// 			defer file.Close()

// 			writer := csv.NewWriter(file)
// 			defer writer.Flush()
// 			productName = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")

// 			if productName == "" {
// 				// If we can't get any name, we return and go directly to the next element
// 				return
// 			}

// 			rating = e.ChildText("span.a-icon-alt")
// 			numberOfRatings = e.ChildText("span.a-size-base.puis-light-weight-text.s-link-centralized-style")

// 			price = e.ChildText("span.a-offscreen")

// 			writer.Write([]string{
// 				productName,
// 				rating,
// 				numberOfRatings,
// 				price,
// 			})
// 		})

// 	})

// 	c.Visit("https://www.amazon.com/s?k=iphone+13&ref=nb_sb_noss_1")
// }
