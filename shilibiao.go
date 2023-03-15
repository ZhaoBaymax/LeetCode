package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	test()
}

func test() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, 0)
	for i := 0; i < 4; i++ {
		scanner.Scan()
		count, _ := strconv.Atoi(scanner.Text())
		for j := 0; j < count; j++ {
			nums = append(nums, i)
		}

	}
	ans := 0
	used := make([]bool, n*n)
	var back func(nums []int, path []int)
	back = func(nums []int, path []int) {
		if len(path) == n*n {
			ans++
			return
		}
		for i := 0; i < n*n; i++ {
			if !used[i] {
				if i >= 1 && nums[i] == nums[i-1] && !used[i-1] {
					continue
				}
				path = append(path, nums[i])
				used[i] = true
				back(nums, path)
				path = path[:len(path)-1]
				used[i] = false
			}

		}
	}
	back(nums, []int{})
	fmt.Println(ans % 998244353)
}
