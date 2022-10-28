package main

import (
	"fmt"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	var nums []int
	for i := 0; i < len(grid); i++ {
		nums = append(nums, grid[i]...)
	}
	sort.Ints(nums)
	mid := len(nums) / 2
	mid_num := nums[mid]
	ret := 0
	for _, v := range nums {
		if abs(v-mid_num)%x != 0 {
			return -1
		}
		ret = ret + abs(v-mid_num)/x
	}
	return ret
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	grid := [][]int{{2, 4}, {6, 8}}
	x := 2
	fmt.Println(minOperations(grid, x))

}
