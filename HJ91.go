package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var Gpath func(n, m int) int
	Gpath = func(n, m int) int {
		dp := make([][]int, n+1)
		for i := 0; i <= n; i++ {
			dp[i] = make([]int, m+1)
			dp[i][1] = 1
		}
		for j := 0; j <= m; j++ {
			dp[1][j] = 1
		}
		for i := 2; i <= n; i++ {
			for j := 2; j <= m; j++ {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
		return dp[n][m]
	}
	fmt.Println(Gpath(n+1, m+1))
}
