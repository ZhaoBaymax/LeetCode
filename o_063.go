package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min_price := prices[0]
	max_profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min_price {
			min_price = prices[i]
			continue
		}
		cur_profit := prices[i] - min_price
		max_profit = max(max_profit, cur_profit)
	}
	return max_profit
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	prices := []int{}
	fmt.Println(maxProfit(prices))
}
