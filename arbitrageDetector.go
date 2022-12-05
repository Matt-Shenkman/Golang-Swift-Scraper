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
			var links []string
			game.ForEach(("div.d-flex.flex-column.odds-row.position-relative"), func(_ int, oddse *colly.HTMLElement) {
				if i < 2 {
					i += 1
				} else {
					bestOdds = make([]string, 0)
					links = make([]string, 0)
					oddse.ForEach(("div.d-flex.flex-row.pr-2.pr-lg-0.px-1"), func(_ int, overUnder *colly.HTMLElement) {
						overUnder.ForEach("a.text-decoration-none", func(_ int, overUnder *colly.HTMLElement) {
							temp := overUnder.ChildText("div.font-weight-bold.pt-3.regular-text.text-center")
							temp2 := strings.Fields(temp)
							if temp != "" {
								bestOdds = append(bestOdds, strings.Join(temp2, " "))
							}

						})

					})

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
