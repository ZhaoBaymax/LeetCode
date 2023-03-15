package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	test(n, m)
}
func test(n, m int) {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[0][j] = 1

		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = dp[i][j]%(1e9+7) + dp[k][j-1]*dp[i-k-1][j-1]%(1e9+7)
			}
		}
	}
	fmt.Println(dp[n][m] % (1e9 + 7))
}
