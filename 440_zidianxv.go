package main

import "fmt"

func findKthNumber(n int, k int) int {
	l := 1
	prefix := 1
	for l < k {
		cnt := getCount(prefix, n)
		if l+cnt > k { // 往下走
			prefix *= 10
			l++
		} else { // 往右走
			prefix++
			l += cnt
		}
	}
	return prefix
}

func getCount(prefix, n int) int {
	count := 0
	start, end := prefix, prefix+1
	for start <= n {
		count += min(n+1, end) - start
		start *= 10
		end *= 10
	}
	return count
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(findKthNumber(10, 3))
}
