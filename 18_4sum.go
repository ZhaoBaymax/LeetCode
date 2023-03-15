package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		three_sum := threeSum(nums, i+1, target-nums[i])
		for _, four_sum := range three_sum {
			four_sum = append(four_sum, nums[i])
			ans = append(ans, four_sum)
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func twoSum(nums []int, start int, target int) [][]int {
	left := start
	right := len(nums) - 1
	ans := make([][]int, 0)
	for left < right {
		sum := nums[left] + nums[right]
		l, r := nums[left], nums[right]
		if sum < target {
			for left < right && l == nums[left] {
				left++
			}
		} else if sum > target {
			for left < right && r == nums[right] {
				right--
			}
		} else {
			ans = append(ans, []int{nums[left], nums[right]})
			for left < right && l == nums[left] {
				left++
			}
			for left < right && r == nums[right] {
				right--
			}
		}
	}
	return ans
}

func threeSum(nums []int, start int, target int) [][]int {
	ans := make([][]int, 0)
	for i := start; i < len(nums); i++ {
		two_sum := twoSum(nums, i+1, target-nums[i])
		for _, three_sum := range two_sum {
			three_sum = append(three_sum, nums[i])
			ans = append(ans, three_sum)
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func main() {
	nums := []int{2, 2, 2, 2, 2}
	fmt.Println(fourSum(nums, 8))
}
