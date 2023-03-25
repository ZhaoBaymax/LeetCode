package main

import (
	"fmt"
	"math"
	"sort"
)

func zhongweishu(nums []int) int {
	sort.Ints(nums)
	mid := len(nums) / 2
	sum := getSum(nums, mid)
	return sum
}
func getSum(nums []int, target int) int {
	var sum float64
	for i := 0; i < len(nums); i++ {
		sum += math.Abs(float64(nums[target] - nums[i]))
	}
	return int(sum)
}

func main() {
	var a int
	fmt.Scan(&a)
	nums := make([]int, 0)
	for i := 0; i < a; i++ {
		var ele1, ele2 int
		fmt.Scan(&ele1, &ele2)
		nums = append(nums, ele1)
	}
	fmt.Print(zhongweishu(nums))
}
