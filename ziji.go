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
	T, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < T; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		x := make([]int, n)
		y := make([]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			x[j], _ = strconv.Atoi(scanner.Text())
		}
		for j := 0; j < n; j++ {
			scanner.Scan()
			y[j], _ = strconv.Atoi(scanner.Text())
		}
		fmt.Println(ziji(x, y))
	}
}

func ziji(x, y []int) int {
	//sort.Ints(x)
	sort.Slice(y, func(i, j int) bool {
		if x[i] == x[j] {
			return y[i] > y[j]
		}
		return x[i] < x[j]
	})
	return lengthOfLIS(y)
}
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
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
