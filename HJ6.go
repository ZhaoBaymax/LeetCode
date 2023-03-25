package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	ans := test(n)
	for i := 0; i < len(ans); i++ {
		fmt.Print(ans[i])
		fmt.Print(" ")
	}
}
func test(num int) []int {
	ans := make([]int, 0)
	n := int(math.Sqrt(float64(num)))
	for i := 2; i < n+1; i++ {
		for num%i == 0 {
			ans = append(ans, i)
			num = num / i
		}
	}
	if num != 1 {
		ans = append(ans, num)
	}
	return ans
}
