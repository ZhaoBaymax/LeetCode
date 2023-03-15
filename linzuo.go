package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int
	fmt.Scan(&a)
	nums := make([]int, 0)
	for i := 0; i < a; i++ {
		var ele int
		fmt.Scan(&ele)
		nums = append(nums, ele)
	}
	sum, sum_list := linzuo(nums)
	fmt.Println(sum)
	for i := 0; i < a; i++ {
		fmt.Print(sum_list[i])
		fmt.Print(" ")
	}
}

func linzuo(nums []int) (int, []int) {
	sum := 0
	sum_list := make([]int, 0)
	left, right := 1, len(nums)-1
	sort.Ints(nums)
	sum_list = append(sum_list, nums[0])
	for left < right {
		sum_list = append(sum_list, nums[right])
		sum += abs(sum_list[len(sum_list)-1], sum_list[len(sum_list)-2])
		sum_list = append(sum_list, nums[left])
		sum += abs(sum_list[len(sum_list)-1], sum_list[len(sum_list)-2])
		right--
		left++
	}
	if left == right {
		sum_list = append(sum_list, nums[right])
		sum += abs(sum_list[len(sum_list)-1], sum_list[len(sum_list)-2])
		sum += abs(sum_list[len(sum_list)-1], sum_list[0])
	}
	if left > right {
		sum += abs(sum_list[len(sum_list)-1], sum_list[0])
	}
	return sum, sum_list
}
func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
