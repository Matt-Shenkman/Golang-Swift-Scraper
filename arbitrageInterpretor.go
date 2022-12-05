package main

import (
	"fmt"
	"math"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func avg(nums []int) float64 {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))
}

func findSTD(nums []int, m float64) float64 {
	sum := float64(0)
	for _, n := range nums {
		sum += math.Pow(math.Abs(float64(n)-m), 2)
	}
	return math.Pow((sum / float64(len(nums))), 0.5)
}

func detectArbitrage(games []GameData) {
	ad := false
	for _, game := range games {
		var unders []int
		var overs []int

		if len(game.t1Odds) == 0 || len(game.t2Odds) == 0 {
			continue
		}
		if game.t1Odds[0] < 0 {
			unders = game.t1Odds
			overs = game.t2Odds
		} else {
			unders = game.t2Odds
			overs = game.t1Odds
		}

		maxminelt := -100000
		mean := avg(unders)
		std := findSTD(unders, mean)
		threshold := mean + float64(2)*std
		for _, n := range unders {
			if float64(n) < threshold {
				maxminelt = max(maxminelt, n)
			}
		}

		maxelt := 0
		mean2 := avg(overs)
		std2 := findSTD(overs, mean2)
		threshold2 := mean2 + float64(2)*std2
		for _, n := range overs {

			if float64(n) < threshold2 {
				maxelt = max(maxelt, n)
			}
		}

		if math.Abs(float64(maxminelt)) < float64(maxelt) && game.date != "Final" {
			ad = true
			fmt.Println("")
			fmt.Println("Arbitrage Detected!")
			fmt.Println(game.team1 + " VS " + game.team2 + " On " + game.date)
			fmt.Println("Over:" + strconv.Itoa(maxelt))
			fmt.Println("Under:" + strconv.Itoa(maxminelt))
		}
	}
	if !ad {
		fmt.Println("No Arbtrage Detected")
	}
}

func main() {
	detectArbitrage(createGameDataTable())
}
