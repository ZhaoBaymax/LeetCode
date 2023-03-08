package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if mid%2 != 0 {
			if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			if nums[mid] == nums[mid+1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return nums[left]
}

func main() {
	arr := []int{1, 1, 2, 2, 3, 3, 4, 5, 5}
	fmt.Println(singleNonDuplicate(arr))
}
