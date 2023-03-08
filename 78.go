package main

import "fmt"

func subsets(nums []int) [][]int {
	track := make([]int, 0)
	var ans [][]int
	var backtrack func(nums []int, start int, track []int)
	backtrack = func(nums []int, start int, track []int) {
		temp := make([]int, len(track))
		copy(temp, track)
		ans = append(ans, temp)
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(nums, i+1, track)
			track = track[:len(track)-1]
		}
	}
	backtrack(nums, 0, track)
	return ans
}
func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
