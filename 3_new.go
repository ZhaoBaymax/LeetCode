package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	win := map[byte]int{}
	count := 0
	left, right := 0, 0
	for right < len(s) {
		win[s[right]]++
		for win[s[right]] > 1 {
			win[s[left]]--
			left++

		}
		right++
		count = max(count, right-left)
	}
	return count
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
