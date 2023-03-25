package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	n := make([]int, x)
	for i := 0; i < x; i++ {
		scanner.Scan()
		n[i], _ = strconv.Atoi(scanner.Text())
	}
	used := make([]bool, x)
	var dfs func(nums []int, path []int)
	dfs = func(nums []int, path []int) {
		if len(path) == len(nums) {
			for _, v := range path {
				fmt.Print(nums[v])
				fmt.Print(" ")
			}
			fmt.Println()
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				if i >= 1 && nums[i] == nums[i-1] && !used[i-1] {
					continue
				}
				path = append(path, i)
				flag := false
				for j := 0; j < len(path); j++ {
					if lengthOfLIS(path, j) > 1 {
						path = path[:len(path)-1]
						flag = true
						break
					}
				}
				if flag {
					continue
				}
				used[i] = true
				dfs(nums, path)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(n, []int{})
}
func lengthOfLIS(nums []int, start int) int {
	dp := make([]int, len(nums))
	for i := start; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := start; i < len(nums); i++ {
		if nums[i] > nums[start] {
			continue
		}
		for j := start; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	sort.Ints(dp)
	return dp[len(dp)-1]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
