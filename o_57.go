package main

import "fmt"

func findContinuousSequence(target int) (ans [][]int) {
	nums := make([]int, target/2+2)
	for i := 0; i <= target/2+1; i++ {
		nums[i] = i
	}
	sum := 0
	left, right := 1, 1
	for right <= target/2+1 {
		sum += right
		right++
		for sum > target {
			sum = sum - left
			left++
		}
		if sum == target && right-left >= 2 {
			ans = append(ans, nums[left:right])
			sum = sum - left
			left++
		}
	}
	return
}
func main() {
	fmt.Println(findContinuousSequence(9))
}
