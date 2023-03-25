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
	nums := make([]int, T)
	for i := 0; i < T; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(myLIS(nums))
}
func myLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j >= 0; j-- {
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
