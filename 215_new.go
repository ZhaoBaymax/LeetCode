package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	return selectKthLargest(nums, 0, len(nums)-1, k)
}
func selectKthLargest(nums []int, left, right, k int) int {
	q := partition(nums, left, right)
	if q+1 == k {
		return nums[q]
	} else if q+1 < k {
		return selectKthLargest(nums, q+1, right, k)
	} else {
		return selectKthLargest(nums, left, q-1, k)
	}
}
func partition(nums []int, left, right int) int {
	k := left + rand.Intn(right-left+1)
	nums[k], nums[right] = nums[right], nums[k]
	i := left - 1
	for j := left; j < right; j++ {
		if nums[j] >= nums[right] {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}

	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}
func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(arr, 2))
}
