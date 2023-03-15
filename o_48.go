package main

func lengthOfLongestSubstring(s string) int {
	win := map[byte]int{}
	ans := 0
	left, right := 0, 0
	for right < len(s) {
		win[s[right]]++
		for win[s[right]] > 1 {
			win[s[left]]--
			left++
		}
		right++
		ans = max(ans, right-left)
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
