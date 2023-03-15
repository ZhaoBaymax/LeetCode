package main

import (
	"container/list"
	"fmt"
)

type MyQue struct {
	ll *list.List
}

func (q MyQue) getMax() int {
	return q.ll.Front().Value.(int)
}
func (q MyQue) push(x int) {
	for ele := q.ll.Back(); q.ll.Len() > 0 && ele.Value.(int) < x; {
		ele = ele.Prev()
		q.ll.Remove(q.ll.Back())
	}
	q.ll.PushBack(x)
}
func (q MyQue) pop(x int) {
	if x == q.ll.Front().Value.(int) {
		q.ll.Remove(q.ll.Front())
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	win := MyQue{ll: list.New()}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			win.push(nums[i])
		} else {
			win.push(nums[i])
			res = append(res, win.getMax())
			win.pop(nums[i-k+1])
		}
	}
	return res
}

func main() {
	nums := []int{1, -1}
	fmt.Println(maxSlidingWindow(nums, 1))
}
