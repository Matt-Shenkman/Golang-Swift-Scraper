package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type GameData struct {
	team1  string
	team2  string
	date   string
	t1Odds []int
	t2Odds []int
}

func stringToIntArr(a []string) []int {
	output := make([]int, 0)
	for _, s := range a {
		intVar, err := strconv.Atoi(s)
		if err != nil {
			intVar = 0
		}
		output = append(output, intVar)
	}
	return output
}
func createGameDataTable() []GameData {
	c := colly.NewCollector(colly.AllowedDomains())
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	gd := make([]GameData, 0)
	c.OnHTML("div.bc-odds-table.bc-table", func(e *colly.HTMLElement) {
		// var p Product
		var g GameData
		e.ForEach("div.d-flex.flex-row.hide-scrollbar.odds-slider-all.syncscroll.tracks", func(_ int, game *colly.HTMLElement) {

			teams := game.ChildText("div.d-block.d-lg-none")
			ts := strings.Fields(teams)
			g.team1 = ts[0]
			g.team2 = ts[1]

			g.date = game.ChildText("div.ml-2.my-1.py-2.regular-text.text-muted")
			i := 0
			var t1Odds []string
			var t2Odds []string
			game.ForEach(("div.d-flex.flex-column.odds-row.position-relative"), func(_ int, oddse *colly.HTMLElement) {
				if i < 2 {
					i += 1
				} else {
					t1Odds = make([]string, 0)
					t2Odds = make([]string, 0)

					j := 0
					oddse.ForEach(("div.d-flex.flex-row.pr-2.pr-lg-0.px-1"), func(_ int, overUnder *colly.HTMLElement) {
						overUnder.ForEach("a.text-decoration-none", func(_ int, overUnder *colly.HTMLElement) {
							temp := overUnder.ChildText("div.font-weight-bold.pt-3.regular-text.text-center")
							temp2 := strings.Fields(temp)
							if temp != "" {
								if j == 0 {
									t1Odds = append(t1Odds, strings.Join(temp2, " "))
								} else {
									t2Odds = append(t2Odds, strings.Join(temp2, " "))
								}
							}

						})
						j += 1

					})
					g.t1Odds = stringToIntArr(t1Odds)
					g.t2Odds = stringToIntArr(t2Odds)

				}

			})
			gd = append(gd, g)
		})

	})

	c.Visit("https://www.vegasinsider.com/nba/odds/las-vegas/")
	return gd[:len(gd)-1]
}
