package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max_sum := nums[0]
	cur_max := nums[0]
	for i := 1; i < len(nums); i++ {
		cur_max = max(cur_max+nums[i], nums[i])
		max_sum = max(max_sum, cur_max)
	}
	return max_sum
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}
