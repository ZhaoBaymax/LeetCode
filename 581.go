package main

import "math"

func findUnsortedSubarray(nums []int) int {
	max_n, min_n := math.MinInt, math.MaxInt
	left, right := -1, -1
	for k, v := range nums {
		if max_n > v {
			right = k
		} else {
			max_n = v
		}
		if min_n < nums[len(nums)-k-1] {
			left = len(nums) - k - 1
		} else {
			min_n = nums[len(nums)-k-1]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}
