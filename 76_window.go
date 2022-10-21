package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	start := 0
	ans_len := math.MaxInt
	valid := 0
	for right < len(s) {
		if _, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				valid++
			}
		}
		right++
		for valid == len(need) {
			if right-left < ans_len {
				start = left
				ans_len = right - left
			}
			if _, ok := need[s[left]]; ok {
				if need[s[left]] == window[s[left]] {
					valid--
				}
				window[s[left]]--
			}
			left++
		}
	}
	if ans_len == math.MaxInt {
		return ""
	}
	return s[start : start+ans_len]
}

func main() {
	s := "a"
	t := "aa"
	fmt.Println(minWindow(s, t))
}
