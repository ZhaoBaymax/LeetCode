package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	set := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		set[nums[i]]++
	}
	flag := false
	for ; head != nil; head = head.Next {
		if _, ok := set[head.Val]; ok {
			if flag == false {
				ans++
				flag = true
			}
		} else {
			flag = false
		}
	}
	return ans
}
