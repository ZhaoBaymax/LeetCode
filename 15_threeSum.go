package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		twosum := twoSum(nums, i+1, -nums[i])
		for _, val := range twosum {
			res = append(res, append(val, nums[i]))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSum(nums []int, start, target int) [][]int {
	l := start
	r := len(nums) - 1
	ans := make([][]int, 0)
	for l < r {
		sum := nums[l] + nums[r]
		left := nums[l]
		right := nums[r]
		if sum < target {
			for l < r && left == nums[l] {
				l++
			}
		} else if sum > target {
			for l < r && right == nums[r] {
				r--
			}
		} else {
			ans = append(ans, []int{nums[l], nums[r]})
			for l < r && left == nums[l] {
				l++
			}
			for l < r && right == nums[r] {
				r--
			}
		}
	}
	return ans
}
