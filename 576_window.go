package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	s1, s2 = s2, s1
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s2); i++ {
		need[s2[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s1) {
		if _, ok := need[s1[right]]; ok {
			window[s1[right]]++
			if window[s1[right]] == need[s1[right]] {
				valid++
			}
		}
		right++
		for right-left >= len(s2) {
			if valid == len(need) {
				return true
			}
			if _, ok := need[s1[left]]; ok {
				if need[s1[left]] == window[s1[left]] {
					valid--
				}
				window[s1[left]]--
			}
			left++
		}
	}
	return false
}

func main() {
	s1 := "abcdxabcde"
	s2 := "abcdeabcdx"
	fmt.Println(checkInclusion(s1, s2))
}
