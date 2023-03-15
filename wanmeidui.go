package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	var k int
	fmt.Scan(&n, &k)
	nums := make([][]int, 0)
	for i := 0; i < n; i++ {
		row := make([]int, 0)
		for j := 0; j < k; j++ {
			var ele int
			fmt.Scan(&ele)
			row = append(row, ele)
		}
		nums = append(nums, row)
	}
	fmt.Println(wanmeidui(nums))
}
func wanmeidui(nums [][]int) int {
	hash_sub := map[string]int{}
	ans := 0
	n := len(nums)
	m := len(nums[0])
	for i := 0; i < n; i++ {
		var sub string
		var cmp string
		for j := 1; j < m; j++ {
			s := nums[i][j] - nums[i][j-1]
			if s < 0 {
				sub = sub + strconv.Itoa(s)
				cmp = cmp + strings.Replace(strconv.Itoa(s), "-", "+", 1)
			} else if s > 0 {
				sub = sub + "+" + strconv.Itoa(s)
				cmp = cmp + "-" + strconv.Itoa(s)
			} else {
				sub = sub + "+" + strconv.Itoa(s)
				cmp = cmp + "+" + strconv.Itoa(s)
			}
		}
		if _, ok := hash_sub[cmp]; ok {
			ans += hash_sub[cmp]
		}
		hash_sub[sub]++
	}
	return ans
}
