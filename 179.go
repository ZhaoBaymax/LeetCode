package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	nums_string := make([]string, 0)
	for _, num := range nums {
		nums_string = append(nums_string, strconv.Itoa(num))
	}
	sort.Slice(nums_string, func(i, j int) bool {
		return nums_string[i]+nums_string[j] > nums_string[j]+nums_string[i]
	})
	var ans string
	for _, ss := range nums_string {
		ans += ss
	}
	if ans[0] == '0' {
		return "0"
	}
	return ans
}

func main() {
	nums := []int{34323, 3432}
	fmt.Println(largestNumber(nums))
}
