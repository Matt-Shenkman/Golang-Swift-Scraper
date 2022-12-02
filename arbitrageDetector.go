package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains())
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	fName := "names.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c.OnHTML("div.bc-odds-table.bc-table", func(e *colly.HTMLElement) {
		// var p Product
		e.ForEach("div.d-flex.flex-row.hide-scrollbar.odds-slider-all.syncscroll.tracks", func(_ int, game *colly.HTMLElement) {
			teams := game.ChildText("div.d-block.d-lg-none")
			ts := strings.Fields(teams)
			i := 0
			var bestOdds []string
			var numbers []int
			game.ForEach(("div.d-flex.flex-column.odds-row.position-relative"), func(_ int, oddse *colly.HTMLElement) {
				if i < 2 {
					i += 1
				} else {
					j := 0
					numbers = make([]int, 0)
					oddse.ForEach(("div.d-flex.flex-row.pr-2.pr-lg-0.px-1"), func(_ int, siteOdds *colly.HTMLElement) {
						numbers = append(numbers, j)
						j += 1
					})
					temp := oddse.ChildText("div.best-odds-box.m-1.odds-box")
					bestOdds = strings.Fields(temp)
					fmt.Println(numbers)
				}

			})

			writer.Write([]string{
				ts[0],
				ts[1],
				strings.Join(bestOdds, " "),
			})

		})

	})

	c.Visit("https://www.vegasinsider.com/college-basketball/odds/las-vegas/")
}
