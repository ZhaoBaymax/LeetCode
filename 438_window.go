package main

import "fmt"

func findAnagrams(s string, p string) []int {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	var res []int
	for right < len(s) {
		if _, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				valid++
			}
		}
		right++
		if right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
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
	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}
