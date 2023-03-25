package main

import (
	"fmt"
	"strings"
)

func main() {
	//var s1, s2 string
	//fmt.Scan(&s1, &s2)
	//lcs := LCS(s1, s2)
	//fmt.Println(lcs)
	s1 := "010101"
	a := strings.LastIndex(s1, "1")
	fmt.Println(a)
}
func LCS(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
		dp[i][0] = i
	}
	for j := 0; j <= len(s2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	}
	if y <= x && y <= z {
		return y
	}
	return z
}
