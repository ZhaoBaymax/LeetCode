package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(dp(m, n))
}
func dp(m, n int) int {
	if m < n {
		return dp(m, m)
	}
	if m == 0 || n == 1 {
		return 1
	}
	return dp(m-n, n) + dp(m, n-1)
}
