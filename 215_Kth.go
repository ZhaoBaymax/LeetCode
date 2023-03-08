package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	m := len(nums) - k + 1 // mth smallest, from 1..len(nums)
	return selectSmallest(nums, 0, len(nums)-1, m)
}
func selectSmallest(nums []int, left, right, m int) int {
	if left >= right {
		return nums[left]
	}

}
func partition(nums []int, left, right int) int {
	k := left + rand.Intn(right-left+1)
	nums[k], nums[right] = nums[right], nums[k]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] <= nums[right] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}
