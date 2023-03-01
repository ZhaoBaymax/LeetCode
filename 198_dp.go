package main

import "fmt"

func robRange(nums []int, start, end int) int {
	//dp := make([]int, n)
	dp_i_1 := max(nums[start], nums[start+1])
	dp_i_2 := nums[start]
	dp_i := 0
	for i := start + 2; i <= end; i++ {
		dp_i = max(dp_i_1, dp_i_2+nums[i])
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i_1
}
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

//	func dp(nums []int, start int, mem []int) int {
//		if start >= len(nums) {
//			return 0
//		}
//		if mem[start] != -1 {
//			return mem[start]
//		}
//		res := max(dp(nums, start+1, mem), nums[start]+dp(nums, start+2, mem))
//		mem[start] = res
//		return res
//	}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	arr := []int{0, 0}
	fmt.Println(rob(arr))
}
