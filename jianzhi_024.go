package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := &ListNode{
		Val:  head.Val,
		Next: nil,
	}
	head = head.Next
	for head != nil {
		temp := head.Next
		head.Next = res
		res = head
		head = temp
	}
	return res
}
