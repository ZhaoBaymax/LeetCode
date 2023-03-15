package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			var ele int
			fmt.Scan(&ele)
			nums = append(nums, ele)
		}
		fmt.Println(guohe(nums, n))
	}
}
func guohe(nums []int, n int) int {
	ans := 0
	sort.Ints(nums)
	for n > 3 {
		ans += min(nums[n-1]+nums[n-2]+2*nums[0], 2*nums[1]+nums[0]+nums[n-1])
		n = n - 2
	}
	if n == 3 {
		ans += nums[2] + nums[0] + nums[1]
	}
	if n == 2 {
		ans += nums[1]
	}
	if n == 1 {
		ans += nums[0]
	}
	return ans
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
