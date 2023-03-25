package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Print(solver(s))
}
func solver(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	ans := 0
	for l := len(s) - 2; l >= 0; l-- {
		for r := l + 1; r < len(s); r++ {
			if s[l] == s[r] {
				dp[l][r] = dp[l+1][r-1] + 2
				ans = max(ans, dp[l][r])
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
