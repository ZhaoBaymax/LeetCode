package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	hari := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := hari
	p := head
	for p != nil {
		tail := p
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hari.Next
			}
		}
		next := tail.Next
		p, tail = myReverse(p, tail)
		pre.Next = p
		tail.Next = next
		pre = tail
		p = pre.Next
	}
	return hari.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	next := tail.Next
	p := head
	for tail != next {
		temp := p.Next
		p.Next = next
		next = p
		p = temp
	}
	return tail, head
}
