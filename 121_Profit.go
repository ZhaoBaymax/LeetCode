package main

import "fmt"

func maxProfit(prices []int) int {
	maxP := 0
	minP := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minP {
			minP = prices[i]
		}
		curProfit := prices[i] - minP
		if maxP < curProfit {
			maxP = curProfit
		}
	}
	return maxP
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
