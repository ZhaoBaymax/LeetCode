package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}
	asc, desc := LIS(nums)
	member := 0
	for i := 0; i < n; i++ {
		ma := getMax(asc, 0, i)
		mi := getMax(desc, i, len(desc)-1)
		member = max(ma+mi-1, member)
	}
	fmt.Println(n - member)
}

func LIS(nums []int) ([]int, []int) {
	asc := make([]int, len(nums))
	desc := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		asc[i] = 1
		desc[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j >= 0; j-- {
			if nums[i] > nums[j] {
				asc[i] = max(asc[i], asc[j]+1)
			}
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				desc[i] = max(desc[i], desc[j]+1)
			}
		}
	}
	return asc, desc
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getMax(nums []int, start, end int) int {
	m := 0
	for i := start; i <= end; i++ {
		m = max(m, nums[i])
	}
	return m
}
