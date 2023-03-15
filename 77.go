package main

import "fmt"

func combine(n int, k int) [][]int {
	if k <= 0 || n <= 0 {
		return [][]int{}
	}
	var ans [][]int
	var back func(n int, start int, path []int)
	back = func(n int, start int, path []int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			back(n, i+1, path)
			path = path[:len(path)-1]
		}
	}
	back(n, 1, []int{})
	return ans
}
func main() {
	fmt.Println(combine(4, 2))
}
