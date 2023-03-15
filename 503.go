package main

import (
	"container/list"
	"fmt"
)

func nextGreaterElements(nums []int) []int {
	stack := list.New()
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = -1
	}
	nn := make([]int, 0)
	nn = append(nn, nums...)
	nn = append(nn, nums...)
	for k, v := range nn {
		for stack.Len() != 0 && v > nums[stack.Back().Value.(int)%len(nums)] {
			ans[stack.Remove(stack.Back()).(int)%len(nums)] = v
		}
		stack.PushBack(k)
	}
	return ans
}
func main() {
	nums := []int{1, 2, 3, 4, 3}
	fmt.Println(nextGreaterElements(nums))

}
