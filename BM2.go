package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here
	hair := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := hair
	p := head
	for i := 0; i < m-1; i++ {
		pre = pre.Next
		p = p.Next
	}
	tail := p
	for i := 0; i < n-m; i++ {
		tail = tail.Next
	}
	next := tail.Next
	head, tail = myReverse(p, tail)
	pre.Next = head
	tail.Next = next
	return hair.Next
}
func myReverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	nex := tail.Next
	p := head
	for nex != tail {
		temp := p.Next
		p.Next = nex
		nex = p
		p = temp
	}
	return tail, head
}
func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	reverseBetween(head, 2, 4)
}
