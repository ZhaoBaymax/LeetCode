package main

import (
	"fmt"
)

func maxSubarraySumCircular(nums []int) int {
	sum := nums[0]
	last_max := nums[0]
	last_min := nums[0]
	res_max := nums[0]
	res_min := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]

		last_max = max(last_max+nums[i], nums[i])
		res_max = max(res_max, last_max)

		last_min = min(last_min+nums[i], nums[i])
		res_min = min(res_min, last_min)
	}
	if res_max > 0 {
		return max(res_max, sum-res_min)
	}
	return res_max
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums := []int{-3, -2, -3}
	fmt.Println(maxSubarraySumCircular(nums))
}
