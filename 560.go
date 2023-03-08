package main

import "fmt"

func subarraySum(nums []int, k int) int {
	n := len(nums)
	preSum := make(map[int]int, n)
	preSum[0] = 1
	ans := 0
	sum0_1 := 0
	for i := 0; i < n; i++ {
		sum0_1 += nums[i]
		sum0_j := sum0_1 - k
		if _, ok := preSum[sum0_j]; ok {
			ans += preSum[sum0_j]
		}
		preSum[sum0_1]++
	}
	return ans
}
func main() {
	arr := []int{3, 5, 2, -2, 4, 1}
	fmt.Println(subarraySum(arr, 5))
}
