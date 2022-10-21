package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	window := map[byte]int{}

	left, right := 0, 0
	res := 0
	for right < len(s) {
		window[s[right]]++
		for window[s[right]] > 1 {
			window[s[left]]--
			left++
		}
		right++
		res = max(res, right-left)
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
